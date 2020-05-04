package echo

import (
	"context"
	"go.opentelemetry.io/otel/api/global"

	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
	log "github.com/kaito2/rest-api-sample/gen/log"
)

// echo-service service example implementation.
// The example methods log the requests and return zero values.
type echoServicesrvc struct {
	logger *log.Logger
}

// NewEchoService returns the echo-service service implementation.
func NewEchoService(logger *log.Logger) echoservice.Service {
	return &echoServicesrvc{logger}
}

// EchoGet implements echo-get.
func (s *echoServicesrvc) EchoGet(ctx context.Context, p *echoservice.EchoGetPayload) (res string, err error) {
	// Create custom span.
	tracer := global.TraceProvider().Tracer("/echo-get")
	_, span := tracer.Start(ctx, "sample span")
	defer span.End()

	s.logger.Info().Fields(map[string]interface{}{"message": "hoge", "severity": "warn"}).Send()
	return "sample response", nil
}
