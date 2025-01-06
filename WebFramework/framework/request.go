package framework

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Request struct {
	*http.Request
}

func (r *Request) DecodeBodyInto(obj interface{}) error {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(obj); err != nil {
			return err
		}
	case "application/xml", "text/xml":
		if err := xml.NewDecoder(r.Body).Decode(obj); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported Content-Type: %s", contentType)
	}

	return nil
}
