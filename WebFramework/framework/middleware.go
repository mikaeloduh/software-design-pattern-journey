package framework

// Middleware is a function that is called before the handler
type Middleware func(w *ResponseWriter, r *Request, next func()) error

// func UseRequestScopeMiddleware(container *Container) Middleware {
// 	return func(w *ResponseWriter, r *Request, next func()) error {
// 		container.Get("MyRequestScopeService").(*MyRequestScopeService).SetRequest(r)

// 		next()

// 		return nil
// 	}
// }
