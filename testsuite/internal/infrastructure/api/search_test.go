package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-service-boilerplate/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBoilerSearchHandler_Success(t *testing.T) {
	for i := 0; i < 15; i++ {
		message := "Test Boiler " + domain.RandomString(5)
		boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: message}
		boilerRepository.Create(boiler)
	}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers?search=Test&page=1&size=10", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []*domain.Boiler
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 10)
}

func TestBoilerSearchHandler_InvalidPageNumber(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers?search=Test&page=invalid&size=10", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerSearchHandler_InvalidPageSize(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers?search=Test&page=1&size=invalid", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerSearchHandler_EmptyPageNumberItsEqualsToDefaultWithoutErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers?search=Test&page=&size=1", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestBoilerSearchHandler_EmptyPageSizeItsEqualsToDefaultWithoutErrors(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers?search=Test&page=1&size=", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}
