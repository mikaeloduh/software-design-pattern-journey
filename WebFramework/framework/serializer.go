package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ReadBodyAsObject reads the request body and un-marshals it into the provided object.
// It supports both JSON and XML based on the Content-Type header.
func (c *Context) ReadBodyAsObject(obj interface{}) error {
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		return json.NewDecoder(c.Request.Body).Decode(obj)
	case "application/xml", "text/xml":
		return xml.NewDecoder(c.Request.Body).Decode(obj)
	default:
		return fmt.Errorf("unsupported Content-Type: %s", contentType)
	}
}

// WriteObjectAsJSON sets the Content-Type header to "application/json"
// and writes the JSON-encoded object to the response writer.
func (c *Context) WriteJSON(obj interface{}) error {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(c.ResponseWriter).Encode(obj)
}

// WriteObjectAsXML sets the Content-Type header to "application/xml"
// and writes the XML-encoded object to the response writer.
func (c *Context) WriteXML(obj interface{}) error {
	c.ResponseWriter.Header().Set("Content-Type", "application/xml")
	return xml.NewEncoder(c.ResponseWriter).Encode(obj)
}
