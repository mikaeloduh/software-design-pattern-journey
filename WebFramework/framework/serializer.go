package framework

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadBodyAsObject(r *http.Request, reqData interface{}) error {
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(reqData); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("Unsupported Content-Type: %s", contentType)
}

//func ReadBodyAsObject[T any](r *http.Request) (T, error) {
//	var obj T
//	err := json.NewDecoder(r.Body).Decode(&obj)
//	return obj, err
//}

// WriteObjectAsJSON sets the Content-Type header to "application/json"
// and writes the JSON-encoded object to the response writer.
func WriteObjectAsJSON(w http.ResponseWriter, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(obj)
}
