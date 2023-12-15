package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenRequestCorrect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	assert.NotEmpty(t, data)
}

func TestMainHandlerWhenCityNotSupport(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=maskva&count=1", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	require.Nil(t, err)
	require.Equal(t, res.StatusCode, http.StatusBadRequest)

	assert.Equal(t, string(data), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=7", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	require.Nil(t, err)
	require.Equal(t, res.StatusCode, http.StatusOK)
	assert.Equal(t, len(strings.Split(string(data), ",")), len(cafeList["moscow"]))
}
