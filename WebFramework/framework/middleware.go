package framework

import (
	"net/http"
)

type Middleware func(w http.ResponseWriter, r *http.Request, next func()) error
