package contextfactory

import (
	"io"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ContextOptions for the context factory method.
type ContextOptions struct {
	Method      string
	Path        string
	Body        io.Reader
	PathParams  gin.Params
	QueryParams map[string]string
	Headers     map[string]string
	ContextVars map[string]interface{}
}

// BuildGinTestContext creates a test context for Gin framework with the provided
// ContextOptions. It returns a *gin.Context and *httptest.ResponseRecorder, which
// can be used to simulate HTTP requests and test Gin handlers.
func BuildGinTestContext(opt ContextOptions) (*gin.Context, *httptest.ResponseRecorder) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	// Create the request
	context.Request = httptest.NewRequest(opt.Method, opt.Path, opt.Body)

	// Put path params in the request
	context.Params = opt.PathParams

	// Put query params in the request
	values := url.Values{}
	for k, v := range opt.QueryParams {
		values[k] = []string{v}
	}
	context.Request.URL.RawQuery = values.Encode()

	// Put the header in the request
	for header, value := range opt.Headers {
		context.Request.Header.Set(header, value)
	}

	// Put context vars into the context
	for key, value := range opt.ContextVars {
		context.Set(key, value)
	}

	return context, writer
}
