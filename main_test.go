package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingEndpoint(t *testing.T) {
	router := SetupRouter()
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `{"message":"pong"}`
	assert.JSONEq(t, expected, recorder.Body.String())
}
