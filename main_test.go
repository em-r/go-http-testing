package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewIndex(t *testing.T) {
	assert := assert.New(t)
	req, err := http.NewRequest("GET", "", nil)
	assert.NoError(err)

	rec := httptest.NewRecorder()
	viewIndex(rec, req)

	assert.Equal(http.StatusOK, rec.Result().StatusCode)
	assert.Equal("application/json", rec.Header().Get("content-type"))
	assert.Equal(`{
		"message": "Hello Gophers!"
	}`, rec.Body.String())
}

func TestIndex(t *testing.T) {
	handler := router()
	server := httptest.NewServer(handler)
	defer server.Close()

	type testCase struct {
		name, endpoint string
		shouldFail     bool
	}

	tcs := []testCase{
		{name: "TestExistingEndpoint", endpoint: "/home"},
		{name: "TestMissingEndpoint", endpoint: "/badendpoint", shouldFail: true},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			resp, err := http.Get(fmt.Sprintf("%s%s", server.URL, tc.endpoint))
			assert.NoError(err)

			if tc.shouldFail {
				assert.Equal(http.StatusNotFound, resp.StatusCode)
				return
			}

			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.Equal("application/json", resp.Header.Get("content-type"))
		})
	}
}
