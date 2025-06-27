package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"small-app/internal/models"
	"small-app/pkg/ctxmanage"
)

func (h *handler) GetUsers(c *gin.Context) {
	u := h.con.FetchUsers()
	c.JSON(200, u)

}

func (h *handler) Signup(c *gin.Context) {
	// Declare a variable to hold decoded data from request body

	// Get Content-Length from header
	contentLength := c.Request.ContentLength

	// Define your max size (in bytes)
	const maxSize int64 = 1 * 1024 * 1024 // 1MB

	if contentLength > maxSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{
			"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	var newUser models.NewUser

	// Get the traceId from the request. Useful for tracking the request in logs
	traceId := ctxmanage.GetTraceIdOfRequest(c)

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		log.Error().
			Str("TRACE ID", traceId).
			Str("Error", err.Error()).
			Msg("json validation error")
		// Respond with a 400 Bad Request status code and error message
		c.JSON(http.StatusBadRequest, gin.H{
			"error": http.StatusText(http.StatusBadRequest)})

		// Stop further execution
		return
	}

	uData, err := h.con.CreateUser(newUser)
	// If user fetch operation fails, respond with an error
	if err != nil {
		log.Error().
			Str("Trace ID", traceId).
			Str("Error", err.Error()).
			Msg("error in creating the user")
		c.JSON(http.StatusInternalServerError,
			gin.H{"message": "Problem in creating user"})
		return
	}

	c.JSON(http.StatusCreated, uData)
	log.Info().
		Str("TRACE ID", traceId).
		Str("User", newUser.Email).
		Msg("User created successfully")

}
