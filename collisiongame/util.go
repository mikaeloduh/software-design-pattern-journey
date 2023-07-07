package main

import "reflect"

func isSameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func inputCoord() (x1, x2 int) {
	// TODO: implement stdin
	return 0, 0
}
