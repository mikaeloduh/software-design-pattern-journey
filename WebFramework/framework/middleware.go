package framework

import "net/http"

// Middleware is a function that is called before the handler
type Middleware func(w http.ResponseWriter, r *Request, next func()) error
