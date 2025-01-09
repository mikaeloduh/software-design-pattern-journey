package framework

import (
	"encoding/json"
	"encoding/xml"
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
