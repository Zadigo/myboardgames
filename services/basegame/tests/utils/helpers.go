package utils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A helper function to check if the HTTP response has a status code of 200 OK.
func CheckResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	var body []byte
	body, _ = io.ReadAll(recorder.Result().Body)
	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode, string(body))
}
