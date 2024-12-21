package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-svc-boilerplate/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBoilerListHandler_Success(t *testing.T) {
	server := setupTestServer()
	repo := server.Repository().(*domain.BoilerInMemoryRepository)

	boiler1 := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler 1"}
	boiler2 := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler 2"}
	repo.Create(boiler1)
	repo.Create(boiler2)

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/boilers/all", nil)

	server.Router().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []*domain.Boiler
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "Test Boiler 1", response[0].Message)
	assert.Equal(t, "Test Boiler 2", response[1].Message)
}
