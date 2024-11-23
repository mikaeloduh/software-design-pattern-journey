package framework

import (
	"encoding/json"
	"io"
)

func ReadBodyAsObject(r io.ReadCloser, reqData interface{}) error {
	if err := json.NewDecoder(r).Decode(reqData); err != nil {
		return err
	}
	return nil
}

//func ReadBodyAsObject[T any](r *http.Request) (T, error) {
//	var obj T
//	err := json.NewDecoder(r.Body).Decode(&obj)
//	return obj, err
//}
