package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Status check: if status != 200, no need to continue
	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)

	// Check for behaviour when request is correct
	actualResponse := responseRecorder.Body.String()
	elementsCount := strings.Split(actualResponse, ",")
	assert.NotEmpty(t, actualResponse)
	assert.NotEmpty(t, len(elementsCount))
}

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
	assert.Len(t, recievedCount, totalCount)
}

func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=aaaaaaaaa&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Status check: if status != 400, no need to continue
	status := responseRecorder.Code
	require.Equal(t, http.StatusBadRequest, status)

	// Check behaviour if requested city doesn't exist
	wrongCity := "wrong city value"
	resp := responseRecorder.Body.String()

	assert.Equal(t, wrongCity, resp)
}
