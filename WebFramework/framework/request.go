package framework

import (
	"fmt"
	"net/http"
)

type Request struct {
	*http.Request
	BodyParser Decoder
}

// ParseBodyInto decodes the request body into the provided object
func (r *Request) ParseBodyInto(obj interface{}) error {
	if r.BodyParser == nil {
		return fmt.Errorf("body parser not set")
	}

	contentType := r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		return r.BodyParser(r.Body, obj)
	case "application/xml", "text/xml":
		return r.BodyParser(r.Body, obj)
	default:
		return fmt.Errorf("unsupported Content-Type: %s", contentType)
	}
}
