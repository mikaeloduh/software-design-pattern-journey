package framework

import "net/http"

// ScopeContext for retrieving the instance
// For example, HTTP request, WebSocket connection, Session, CLI invocation…
type ScopeContext interface {
	// You can return any type of identifier, such as string, UUID, *http.Request, etc
	ID() any
}

type HttpScopeContext struct {
	req *http.Request
}

func (h *HttpScopeContext) ID() any {
	// You can return the *http.Request as the key
	// or return the request ID (such as session ID, trace ID in the header)
	return h.req
}

// WebSocketScopeContext
type WebSocketScopeContext struct {
	connID string
}

func (w *WebSocketScopeContext) ID() any {
	return w.connID
}
