package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"webframework/errors"
	"webframework/framework"
)

func TestRouter_MethodNotAllowed(t *testing.T) {
	r := framework.NewRouter()
	r.Handle(http.MethodGet, "/", mockHandler)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)

	assert.Equal(t, http.StatusText(http.StatusMethodNotAllowed), w.Body.String())
}

func TestRouter_Custom_NotFound(t *testing.T) {
	r := framework.NewRouter()
	r.HandleError(&framework.JSONErrorHandler{})
	// no routes added

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]string
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusText(http.StatusNotFound), resp["error"])
}

// FakeDbError 在此檔案直接定義，測試用的「模擬 DB 錯誤」
type FakeDbError struct {
	msg string
}

func (f FakeDbError) Error() string {
	return f.msg
}

// MyGenericError 同樣在此定義模擬一般錯誤
type MyGenericError struct {
	msg string
}

func (m MyGenericError) Error() string {
	return m.msg
}

// 這裡直接定義一個攔截器來判斷 FakeDbError (在此檔案宣告的 struct)
func MyDbErrorInterceptor(err error, c *framework.Context, next func()) {
	if e, ok := err.(*errors.Error); ok && e.Err != nil {
		// 用本檔案內定義的 FakeDbError 做型別斷言
		if _, isFakeDbError := e.Err.(FakeDbError); isFakeDbError {
			e.Code = http.StatusInternalServerError
			c.Status(e.Code)
			c.String("Database error occurred!")
			return // 攔截並終止鏈
		}
	}
	next()
}

func MyDefaultErrorCoder(err error, c *framework.Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		if e.Code == 0 {
			e.Code = http.StatusInternalServerError
		}
	}
	next()
}

// 最終 fallback
func FinalFallbackHandler(err error, c *framework.Context, next func()) {
	if e, ok := err.(*errors.Error); ok {
		code := e.Code
		if code == 0 {
			code = http.StatusInternalServerError
		}
		c.Status(code)
		c.String(e.Error())
	} else {
		c.Status(http.StatusInternalServerError)
		c.String(err.Error())
	}
}

// 測試：確保一般錯誤不會被誤判為 DB 錯誤
func TestErrorHandlerChain(t *testing.T) {
	r := framework.NewRouter()

	// 註冊多個 ErrorHandlerFunc
	r.UseErrorHandler(
		MyDbErrorInterceptor, // 先攔截 DB 錯誤
		MyDefaultErrorCoder,  // 將 code=0 改為 500
		FinalFallbackHandler, // fallback
	)

	// 模擬路由 1：產生 DB 錯誤
	r.Handle(http.MethodGet, "/db-error", func(ctx *framework.Context) {
		dbErr := FakeDbError{"db connection timeout"}
		ctx.AbortWithError(errors.NewError(0, dbErr))
	})

	// 模擬路由 2：一般錯誤
	r.Handle(http.MethodGet, "/generic-error", func(ctx *framework.Context) {
		genericErr := MyGenericError{"some generic error"}
		ctx.AbortWithError(errors.NewError(0, genericErr))
	})

	// 模擬路由 3：已指定 code=400
	r.Handle(http.MethodGet, "/bad-request", func(ctx *framework.Context) {
		ctx.AbortWithError(errors.NewError(http.StatusBadRequest, nil))
	})

	t.Run("DB error => 攔截器立刻回應 500", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/db-error", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "Database error occurred!", w.Body.String())
	})

	t.Run("一般錯誤 code=0 => 改成 500 => 最終回應", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/generic-error", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// 預期是 MyDefaultErrorCoder 把 code=0 -> 500，最後由 FinalFallbackHandler 回應
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "some generic error", w.Body.String())
	})

	t.Run("code=400 => 不改 => 最終回應 400", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/bad-request", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, http.StatusText(http.StatusBadRequest), w.Body.String())
	})
}
