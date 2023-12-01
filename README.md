# Context Factory

The Context Factory provides a straightforward method for testing Gin Handlers by generating a test context.

![tests](https://github.com/Aloe-Corporation/contextfactory/actions/workflows/go.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/Aloe-Corporation/contextfactory.svg)](https://pkg.go.dev/github.com/Aloe-Corporation/contextfactory)

## Overview

The module provides:

- A simple method to create `gin.Context`
- An easy way to test your `gin.Handlers`

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