// Code generated by goa v3.1.2, DO NOT EDIT.
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
	v := &echoservice.EchoGetPayload{}
	v.Name = name
	v.Age = age

	return v
}
