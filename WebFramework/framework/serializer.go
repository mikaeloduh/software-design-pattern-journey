package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// ReadBodyAsObject reads the request body and un-marshals it into the provided object.
// It supports both JSON and XML based on the Content-Type header.
func ReadBodyAsObject(r *http.Request, reqData interface{}) error {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
			return err
		}
		return nil
	case "application/xml", "text/xml":
		if err := xml.NewDecoder(r.Body).Decode(reqData); err != nil {
			return err
		}
		return nil
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
