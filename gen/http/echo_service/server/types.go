// Code generated by goa v3.0.10, DO NOT EDIT.
//
// echo-service HTTP server types
//
// Command:
// $ goa gen github.com/kaito2/rest-api-sample/design

package server

import (
	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
)

// NewEchoGetPayload builds a echo-service service echo-get endpoint payload.
func NewEchoGetPayload(name string, age int) *echoservice.EchoGetPayload {
	return &echoservice.EchoGetPayload{
		Name: name,
		Age:  age,
	}
}
