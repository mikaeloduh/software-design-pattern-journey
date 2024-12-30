package framework

import (
	"context"
	"fmt"
	"net/http"
	"webframework/errors"
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
					errorAware.HandleError(errors.NewError(http.StatusInternalServerError, fmt.Errorf("panic recovered")), w, r)
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
