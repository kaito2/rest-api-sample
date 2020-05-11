// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the echo-service service.
//
// Command:
// $ goa gen github.com/kaito2/rest-api-sample/design

package client

import (
	"fmt"
)

// VersionEchoServicePath returns the URL path to the echo-service service version HTTP endpoint.
func VersionEchoServicePath() string {
	return "/version"
}

// EchoGetEchoServicePath returns the URL path to the echo-service service echo-get HTTP endpoint.
func EchoGetEchoServicePath(name string) string {
	return fmt.Sprintf("/name/%v", name)
}
