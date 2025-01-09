package framework

import "net/http"

// JSONBodyParser is a middleware that sets the BodyParser to JSONDecoder
func JSONBodyParser(w http.ResponseWriter, r *Request, next func()) error {
	if r.Header.Get("Content-Type") == "application/json" {
		r.BodyParser = JSONDecoder
	}

	next()

	return nil
}

// XMLBodyParser is a middleware that sets the BodyParser to XMLDecoder
func XMLBodyParser(w http.ResponseWriter, r *Request, next func()) error {
	if r.Header.Get("Content-Type") == "application/xml" {
		r.BodyParser = XMLDecoder
	}

	next()

	return nil
}
