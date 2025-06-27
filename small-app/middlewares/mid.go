package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

type key string

const TraceIdKey key = "1"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestStartTime := time.Now()

		// get a trace id
		traceId := uuid.NewString()

		//fetching the context container from the request.context()
		ctx := c.Request.Context()

		//adding the value in the context
		ctx = context.WithValue(ctx, TraceIdKey, traceId)
		// The 'WithContext' method on 'c.Request' creates a new copy of the request ('req'),
		// but with an updated context ('ctx') that contains our trace ID.
		// The original request does not get changed by this; we're simply creating a new version of it ('req').
		c.Request = c.Request.WithContext(ctx)

		log.Info().
			Str("TRACE ID", traceId).
			Str("Method", c.Request.Method).
			Interface("URL Path", c.Request.URL.Path).
			Msg("started")

		//we use c.Next only when we are using r.Use() method to assign middlewares
		c.Next() // call next thing in the chain

		log.Info().
			Str("TRACE ID", traceId).
			Str("Method", c.Request.Method).
			Interface("URL Path", c.Request.URL.Path).
			Int("Status Code", c.Writer.Status()).
			Int64("duration μs", time.Since(requestStartTime).Microseconds()).
			Msg("completed")
	}
}
