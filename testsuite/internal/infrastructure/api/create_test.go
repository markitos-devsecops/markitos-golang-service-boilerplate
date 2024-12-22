package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-service-boilerplate/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestBoilerCreateHandler_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerCreateRequest{
		Message: "Test Boiler",
	})
	request, _ := http.NewRequest(http.MethodPost, "/boilers", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func TestBoilerCreateHandler_MissingMessage(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerCreateRequest{})
	request, _ := http.NewRequest(http.MethodPost, "/boilers", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	boilerApiServer.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
