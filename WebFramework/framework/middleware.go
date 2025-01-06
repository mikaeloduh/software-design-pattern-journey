package framework

import (
	"net/http"
)

type Middleware func(w http.ResponseWriter, r *Request, next func()) error
