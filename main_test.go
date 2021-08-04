package main

import (
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
