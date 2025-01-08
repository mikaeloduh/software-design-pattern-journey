package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// Decoder is a function that decodes a reader into an object
type Decoder func(io.Reader, interface{}) error

func JSONDecoder(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func XMLDecoder(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

// ReadBodyAsObject reads the request body and un-marshals it into the provided object.
// It supports both JSON and XML based on the Content-Type header.
func ReadBodyAsObject(r *http.Request, obj interface{}) error {
	contentType := r.Header.Get("Content-Type")

	switch contentType {
	case "application/json":
		return JSONDecoder(r.Body, obj)
	case "application/xml", "text/xml":
		return XMLDecoder(r.Body, obj)
	default:
		return fmt.Errorf("unsupported Content-Type: %s", contentType)
	}
}

// WriteObjectAsJSON sets the Content-Type header to "application/json"
// and writes the JSON-encoded object to the response writer.
func WriteObjectAsJSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(obj)
}

// WriteObjectAsXML sets the Content-Type header to "application/xml"
// and writes the XML-encoded object to the response writer.
func WriteObjectAsXML(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/xml")
	return xml.NewEncoder(w).Encode(obj)
}
