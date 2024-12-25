package api_test

import (
	"markitos-golang-service-boilerplate/internal/domain"
	"markitos-golang-service-boilerplate/internal/infrastructure/api"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var boilerApiServer *api.Server
var boilerRepository domain.BoilerRepository

func TestMain(m *testing.M) {
	boilerApiServer = setupTestServer()
	boilerRepository = boilerApiServer.Repository()

	os.Exit(m.Run())
}

func setupTestServer() *api.Server {
	gin.SetMode(gin.TestMode)
	repo := domain.NewBoilerInMemoryRepository()
	return api.NewServer(":8080", repo)
}
