package ctxmanage

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"small-app/middlewares"
)

func GetTraceIdOfRequest(c *gin.Context) string {
	ctx := c.Request.Context()
	// ok is false if the type assertion was not successful
	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("trace id not present in the context")
		traceId = "Unknown"
	}
	return traceId
}
