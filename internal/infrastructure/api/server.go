package api

import (
	"markitos-service-boilerplate/internal/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address    string
	repository domain.BoilerRepository
	router     *gin.Engine
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Repository() domain.BoilerRepository {
	return s.repository
}

func NewServer(address string, repository domain.BoilerRepository) *Server {
	server := &Server{
		address:    address,
		repository: repository,
		router:     nil,
	}
	server.router = server.createRouter()

	return server
}

func (s *Server) createRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/boilers", s.boilerCreateHandler)
	router.GET("/boilers/all", s.boilerListHandler)
	router.GET("/boilers/:id", s.boilerOneHandler)
	router.PUT("/boilers/:id", s.boilerUpdateHandler)
	router.GET("/", s.boilerMotdHandler)
	router.GET("/boilers", s.boilerSearchHandler)

	return router
}

func (s *Server) Run() error {
	return s.router.Run(s.address)
}

func errorResonses(err error) gin.H {
	return gin.H{"error": err.Error()}
}
