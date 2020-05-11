package echo

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"log"
	"time"

	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
	genlog "github.com/kaito2/rest-api-sample/gen/log"
)

// echo-service service example implementation.
// The example methods log the requests and return zero values.
type echoServicesrvc struct {
	tracer trace.Tracer
	logger *genlog.Logger
	projectID string
	apiVersion string
}

// NewEchoService returns the echo-service service implementation.
func NewEchoService(logger *genlog.Logger) echoservice.Service {
	// load env vars
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("failed to envconfig.Process: %v", err)
	}

	// set up Cloud Trace
	tracer := global.TraceProvider().Tracer("/echo-get")

	return &echoServicesrvc{
		logger: logger,
		tracer: tracer,
		projectID: env.GCPProjectID,
		apiVersion: env.APIVersion,
	}
}

func (s *echoServicesrvc) Version(ctx context.Context) (res string, err error) {
	return s.apiVersion, nil
}

// EchoGet implements echo-get.
func (s *echoServicesrvc) EchoGet(ctx context.Context, p *echoservice.EchoGetPayload) (res string, err error) {
	// Create custom span.
	ctx, span := s.tracer.Start(ctx, "sample_span1")
	defer span.End()

	s.sampleFunction(ctx)

	s.logger.Info().Fields(map[string]interface{}{
		"message": "message of sample span1",
		"severity": "info",
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", s.projectID, span.SpanContext().TraceIDString()),
		"logging.googleapis.com/spanId": span.SpanContext().SpanIDString(),
	}).Send()

	time.Sleep(300*time.Millisecond)
	return "sample response2", nil
}

func (s *echoServicesrvc) sampleFunction(ctx context.Context) {
	_, span := s.tracer.Start(ctx, "sample_span2")
	defer span.End()
	s.logger.Info().Fields(map[string]interface{}{
		"message": "message of sample span2",
		"severity": "info",
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", s.projectID, span.SpanContext().TraceIDString()),
		"logging.googleapis.com/spanId": span.SpanContext().SpanIDString(),
	}).Send()
	time.Sleep(500*time.Millisecond)
}
