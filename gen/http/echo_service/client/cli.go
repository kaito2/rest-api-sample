// Code generated by goa v3.1.2, DO NOT EDIT.
//
// echo-service HTTP client CLI support package
//
// Command:
// $ goa gen github.com/kaito2/rest-api-sample/design

package client

import (
	"fmt"
	"strconv"

	echoservice "github.com/kaito2/rest-api-sample/gen/echo_service"
)

// BuildEchoGetPayload builds the payload for the echo-service echo-get
// endpoint from CLI flags.
func BuildEchoGetPayload(echoServiceEchoGetName string, echoServiceEchoGetAge string) (*echoservice.EchoGetPayload, error) {
	var err error
	var name string
	{
		name = echoServiceEchoGetName
	}
	var age int
	{
		var v int64
		v, err = strconv.ParseInt(echoServiceEchoGetAge, 10, 64)
		age = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for age, must be INT")
		}
	}
	v := &echoservice.EchoGetPayload{}
	v.Name = name
	v.Age = age

	return v, nil
}
