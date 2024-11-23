package framework

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Requestx struct {
	*http.Request
}

func (r *Requestx) ReadBodyAsObject(objType interface{}) (interface{}, error) {
	objValue := reflect.New(reflect.TypeOf(objType)).Interface()
	err := json.NewDecoder(r.Body).Decode(objValue)
	if err != nil {
		return nil, err
	}
	return objValue, nil
}
