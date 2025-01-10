package framework

// JSONBodyParser is a middleware that sets the BodyParser to JSONDecoder
func JSONBodyParser(w *ResponseWriter, r *Request, next func()) error {
	if r.Header.Get("Content-Type") == "application/json" {
		r.BodyParser = JSONDecoder
	}

	next()

	return nil
}

// XMLBodyParser is a middleware that sets the BodyParser to XMLDecoder
func XMLBodyParser(w *ResponseWriter, r *Request, next func()) error {
	if r.Header.Get("Content-Type") == "application/xml" {
		r.BodyParser = XMLDecoder
	}

	next()

	return nil
}

func JSONBodyEncoder(w *ResponseWriter, r *Request, next func()) error {
	w.UseEncoder(JSONEncoder)

	accept := r.Header.Get("Accept")
	if accept == "" || accept == "*/*" || accept == "application/json" {
		w.Header().Set("Content-Type", "application/json")
	}

	next()

	return nil
}

func XMLBodyEncoder(w *ResponseWriter, r *Request, next func()) error {
	w.UseEncoder(XMLEncoder)

	if r.Header.Get("Accept") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
	}

	next()

	return nil
}
