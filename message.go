package restlib

import "net/http"

type Message struct {
	// Method tells which method to use when sending request
	Method string

	// Status is a representation of eg. http status code in responses
	Status int

	// Header is a map of Message headers
	Header http.Header

	// Body is the message body object
	// the message body is de-serialized at retrieval and serialized before sending request
	Body interface{}
}

