package commons

import "reflect"

func IsSameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func InputCoord() (x1, x2 int) {
	// TODO: implement stdin
	return 0, 0
}
