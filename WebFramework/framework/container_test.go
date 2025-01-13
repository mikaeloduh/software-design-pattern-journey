package framework

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestService struct {
	value int
}

func NewTestService() *TestService {
	return &TestService{value: 1}
}

func TestContainer_Singleton(t *testing.T) {
	t.Parallel()

	container := NewContainer()

	container.Register("TestService", func() any { return NewTestService() }, &SingletonStrategy{})

	testInstance := container.Get("TestService").(*TestService)
	expectedInstance := container.Get("TestService").(*TestService)

	assert.Same(t, expectedInstance, testInstance)
}

func TestContainer_Singleton_Parallel(t *testing.T) {
	container := NewContainer()
	container.Register("TestService", func() any { return NewTestService() }, &SingletonStrategy{})

	const concurrency = 100
	var wg sync.WaitGroup

	results := make([]*TestService, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			results[i] = container.Get("TestService").(*TestService)
		}(i)
	}

	wg.Wait()

	for i := 0; i < concurrency; i++ {
		assert.Same(t, results[0], results[i])
	}
}

func TestContainer_Prototype(t *testing.T) {
	t.Parallel()

	container := NewContainer()

	container.Register("TestService", func() any { return NewTestService() }, &PrototypeStrategy{})

	service1 := container.Get("TestService").(*TestService)
	service2 := container.Get("TestService").(*TestService)
	assert.NotSame(t, service1, service2)
}

func TestHttpRequestScope_SameRequest(t *testing.T) {
	container := NewContainer()
	container.Register("TestService", func() any { return NewTestService() }, &HttpRequestScopeStrategy{})

	router := NewRouter()
	router.Use(HttpRequestScopeMiddleware(container))
	router.Handle("/test-scope", http.MethodGet, HandlerFunc(func(w *ResponseWriter, r *Request) error {
		service1 := container.GetFromRequest(r.Request, "TestService").(*TestService)
		service2 := container.GetFromRequest(r.Request, "TestService").(*TestService)

		w.Write([]byte(fmt.Sprintf("%p %p", service1, service2)))
		return nil
	}))

	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL + "/test-scope")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// body contains two service addresses, for example "0xc000010000 0xc000010000"
	// confirm that they are the same memory address
	splitBody := strings.Split(string(body), " ")
	assert.Equal(t, splitBody[0], splitBody[1])
}

func TestHttpRequestScope_DifferentRequests(t *testing.T) {
	container := NewContainer()
	container.Register("TestService", func() any { return NewTestService() }, &HttpRequestScopeStrategy{})

	router := NewRouter()
	router.Use(HttpRequestScopeMiddleware(container))
	router.Handle("/test-scope", http.MethodGet, HandlerFunc(func(w *ResponseWriter, r *Request) error {
		service1 := container.GetFromRequest(r.Request, "TestService").(*TestService)

		w.Write([]byte(fmt.Sprintf("%p", service1)))
		return nil
	}))

	server := httptest.NewServer(router)
	defer server.Close()

	resp1, err := http.Get(server.URL + "/test-scope")
	if err != nil {
		t.Fatal(err)
	}
	defer resp1.Body.Close()

	resp2, err := http.Get(server.URL + "/test-scope")
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()

	body1, _ := io.ReadAll(resp1.Body)
	body2, _ := io.ReadAll(resp2.Body)
	// body contains two service addresses
	// confirm that they are not the same memory address
	assert.NotEqual(t, string(body1), string(body2))
}
