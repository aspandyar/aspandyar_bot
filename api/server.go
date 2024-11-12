package api

import (
	"github.com/aspandyar/aspandyar_bot/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for the API.
type Server struct {
	config util.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	server.setupRoutes()
	return server, nil
}

func (server *Server) setupRoutes() {
	router := gin.Default()

	router.GET("/ping", server.handlePing)
	router.GET("/health", server.handleHealth)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
