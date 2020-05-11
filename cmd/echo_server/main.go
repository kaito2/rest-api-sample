package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	echo "github.com/kaito2/rest-api-sample"
	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
	log "github.com/kaito2/rest-api-sample/gen/log"
	"github.com/rs/zerolog"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel/api/global"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logLevel := zerolog.InfoLevel
		zerolog.SetGlobalLevel(logLevel)
		l := zerolog.New(os.Stdout).With().Str("service", echoservice.ServiceName).Logger()
		logger = &log.Logger{Logger: &l}
	}

	// Initialize the services.
	var (
		echoServiceSvc echoservice.Service
	)
	{
		echoServiceSvc = echo.NewEchoService(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		echoServiceEndpoints *echoservice.Endpoints
	)
	{
		echoServiceEndpoints = echoservice.NewEndpoints(echoServiceSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Set up Cloud Trace
	// More details: https://cloud.google.com/trace/docs/setup/go-ot

	// Create exporter.
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.NewExporter(texporter.WithProjectID(projectID))
	if err != nil {
		logger.Fatal().Msgf("texporter.NewExporter (project_id '%s'): %v", projectID, err)
	}

	// Create trace provider with the exporter.
	//
	// By default it uses AlwaysSample() which samples all traces.
	// In a production environment or high QPS setup please use
	// ProbabilitySampler set at the desired probability.
	// Example:
	//   config := sdktrace.Config{DefaultSampler:sdktrace.ProbabilitySampler(0.0001)}
	//   tp, err := sdktrace.NewProvider(sdktrace.WithConfig(config), ...)
	tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
	if err != nil {
		logger.Fatal().Err(err)
	}
	global.SetTraceProvider(tp)

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://0.0.0.0:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, echoServiceEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Info().Msgf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info().Msg("exited")
}
