package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (server *Server) handleHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
