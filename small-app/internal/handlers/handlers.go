package handlers

import (
	"github.com/gin-gonic/gin"
	"small-app/internal/models"
	"small-app/middlewares"
)

type handler struct {
	con models.Conn
}

func API(conn models.Conn) *gin.Engine {
	r := gin.New()

	r.Use(middlewares.Logger(), gin.Recovery())
	setupRoutes(r, conn)

	return r
}

func setupRoutes(r *gin.Engine, conn models.Conn) {

	h := handler{
		con: conn,
	}
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", h.GetUsers)
	r.POST("/create", h.Signup)

}
