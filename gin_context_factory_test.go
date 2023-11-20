package contextfactory

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var buildGinTestContextTestCases = []ContextOptions{
	{ // All field except for the body.
		Method: "GET",
		Path:   "/test",
		Body:   nil,
		PathParams: gin.Params{
			gin.Param{
				Key:   "param_test_key",
				Value: "param_test_value",
			},
		},
		QueryParams: map[string]string{
			"query_test_key": "query_test_value",
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		ContextVar: map[string]interface{}{
			"test": "test",
		},
	},
	{ // All field with the body.
		Method: "GET",
		Path:   "/test",
		Body:   strings.NewReader("test body data"),
		PathParams: gin.Params{
			gin.Param{
				Key:   "param_test_key",
				Value: "param_test_value",
			},
		},
		QueryParams: map[string]string{
			"query_test_key": "query_test_value",
		},
		Headers: map[string]string{
			"header_test_key": "header_test_value",
		},
		ContextVar: map[string]interface{}{
			"test": "test",
		},
	},
}

func TestBuildGinTestContext(t *testing.T) {
	for i, testCase := range buildGinTestContextTestCases {
		t.Run("Case "+strconv.Itoa(i), func(t *testing.T) {
			c, recorder := BuildGinTestContext(testCase)
			assert.NotNil(t, recorder)

			req := c.Request

			assert.Equal(t, testCase.Method, req.Method)
			assert.Equal(t, testCase.Path, req.URL.Path)

			if testCase.Body == nil {
				assert.Equal(t, http.NoBody, req.Body)
			}

			for header_key, header_value := range testCase.Headers {
				assert.Equal(t, header_value, req.Header.Get(header_key))
			}
			for query_key, query_value := range testCase.QueryParams {
				assert.Equal(t, query_value, req.URL.Query()[query_key][0])
			}
			for var_key, var_value := range testCase.ContextVar {
				assert.Equal(t, var_value, c.GetString(var_key))
			}
		})
	}
}
