package echo

import (
	"context"

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
	s.logger.Info().Msg("echoService.echo-get")
	return "sample response", nil
}
