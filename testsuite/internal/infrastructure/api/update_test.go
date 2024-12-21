package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-svc-boilerplate/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBoilerUpdateHandler_Success(t *testing.T) {
	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	boilerRepository.Create(boiler)

	updatedMessage := "Updated Boiler"
	requestBody, _ := json.Marshal(map[string]string{
		"message": updatedMessage,
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/boilers/"+boiler.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response domain.Boiler
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, boiler.Id, response.Id)
	assert.Equal(t, updatedMessage, response.Message)
}

func TestBoilerUpdateHandler_InvalidID(t *testing.T) {
	requestBody, _ := json.Marshal(map[string]string{
		"message": "Updated Boiler",
	})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/boilers/invalid-id", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerUpdateHandler_MissingMessage(t *testing.T) {
	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	boilerRepository.Create(boiler)

	requestBody, _ := json.Marshal(map[string]string{})
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/boilers/"+boiler.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
