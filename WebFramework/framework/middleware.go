package framework

import (
	"context"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// ErrorAwareMiddleware 注入錯誤處理能力到請求上下文中
func ErrorAwareMiddleware(errorAware ErrorAware) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "errorAware", errorAware)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				if errorAware, ok := r.Context().Value("errorAware").(ErrorAware); ok {
					errorAware.HandleError(NewError(ErrorTypeInternalServerError, "panic recovered", nil), w, r)
				} else {
					http.Error(w, "internal server error", http.StatusInternalServerError)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func CustomRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("custom error content"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func LoggerMiddleware(c *Context) {
	// Just a demo: you might log the request here
	// proceed
	c.Next()
}

func RecoveryMiddleware(c *Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithError(http.ErrAbortHandler)
		}
	}()
	c.Next()
}
