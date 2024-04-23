package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Status check: if status != 200, no need to continue
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)

	// Check behaviour if requested count exceeds total count
	recievedCount := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(recievedCount), totalCount)
}
