# Context Factory

The Context Factory provides a straightforward method for testing Gin Handlers by generating a test context.

This project has been developped by the [Aloe](https://www.aloe-corp.com/) team and is now open source.

![tests](https://github.com/Aloe-Corporation/contextfactory/actions/workflows/go.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/Aloe-Corporation/contextfactory.svg)](https://pkg.go.dev/github.com/Aloe-Corporation/contextfactory)

## Overview

The module provides:

- A simple method to create `gin.Context`
- An easy way to test your `gin.Handlers`

## Concept

Emulate an API call by populating a `gin.Context` thanks to the `ContextOptions` structure:

```go
type ContextOptions struct {
	Method      string
	Path        string
	Body        io.Reader
	PathParams  gin.Params
	QueryParams map[string]string
	Headers     map[string]string
	ContextVars map[string]interface{}
}
```

## Usage

Use the factory method to create a `gin.Context` and a `httptest.ResponseRecorder` inside a unit test:

```go
    func TestGetDog(t *testing.T) {
        options := contextfactory.ContextOptions{
            Path: "/kid",
            QueryParams: map[string]string{
                "age":      3,
            }
        }

        context, writer := contextfactory.BuildGinTestContext(testCase.Context)

        GetDogs(context)

        assert.Equal(t, http.StatusOK, writer.Code)
    }
```

## Contributing

This section will be added soon.

## License

Client is released under the MIT license. See [LICENSE.txt](./LICENSE).