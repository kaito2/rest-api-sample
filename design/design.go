package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zerologger"
)

// API describes the global properties of the API server.
var _ = API("echo", func() {
	Title("Echo Service")
	Description("This is HTTP echo service")
	Server("echo-server", func() {
		Host("localhost", func() { URI("http://0.0.0.0:8080") })
	})
})

// Service describes a service
var _ = Service("echo-service", func() {
	Description("Echo your request")
	// Method describes a service method (endpoint)
	Method("echo-get", func() {
		// define request payload
		Payload(func() {
			// Attribute describes an object field
			Attribute("name", String, "Your name")
			Attribute("age", Int, "Your age")
			// Both attributes must be provided when invoking "add"
			Required("name", "age")
		})
		// define response data type
		Result(String)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			GET("/name/{name}")
			Param("age")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
