package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func recorderAndHandler(req *http.Request) (responseRecorder *httptest.ResponseRecorder) {

	responseRecorder = httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	responseRecorder := recorderAndHandler(req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, req.Body, "wrong count value")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

}

func TestMainHandlerWrongCity(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=wrongcity", nil)
	responseRecorder := recorderAndHandler(req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
}
