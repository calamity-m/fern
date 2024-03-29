package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Timer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		dur := time.Since(start)
		log.Printf("Request: Duration - %v | Status - %d | Path - %s", dur, ctx.Writer.Status(), ctx.Request.RequestURI)
	}
}
