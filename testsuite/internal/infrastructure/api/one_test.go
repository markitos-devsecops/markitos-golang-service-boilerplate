package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-service-boilerplate/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBoilerOneHandler_Success(t *testing.T) {
	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	boilerRepository.Create(boiler)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers/"+boiler.Id, nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response domain.Boiler
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, boiler.Id, response.Id)
	assert.Equal(t, boiler.Message, response.Message)
}

func TestBoilerOneHandler_NotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers/non-existent-id", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerOneHandler_InvalidID(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers/invalid-id", nil)

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
