package echo

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/trace"
	"time"

	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
	log "github.com/kaito2/rest-api-sample/gen/log"
)

// echo-service service example implementation.
// The example methods log the requests and return zero values.
type echoServicesrvc struct {
	tracer trace.Tracer
	logger *log.Logger
	projectID string
}

// NewEchoService returns the echo-service service implementation.
func NewEchoService(logger *log.Logger) echoservice.Service {
	projectID := "kaito2"
	tracer := global.TraceProvider().Tracer("/echo-get")
	return &echoServicesrvc{logger: logger, tracer: tracer, projectID: projectID}
}

// EchoGet implements echo-get.
func (s *echoServicesrvc) EchoGet(ctx context.Context, p *echoservice.EchoGetPayload) (res string, err error) {
	// Create custom span.
	_, span := s.tracer.Start(ctx, "sample_span1")
	defer span.End()

	s.logger.Info().Fields(map[string]interface{}{
		"message": "message of sample span1",
		"severity": "info",
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", s.projectID, span.SpanContext().TraceIDString()),
		"logging.googleapis.com/spanId": span.SpanContext().SpanIDString(),
	}).Send()

	time.Sleep(300*time.Millisecond)
	return "sample response", nil
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
