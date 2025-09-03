package middleware

import (
	"github.com/gin-gonic/gin"
)

func Next(ctx *gin.Context) {
	ctx.Next()
}

func Block(ctx *gin.Context) {
	ctx.JSON(401, map[string]any{
		"err": "401",
	})
	ctx.Abort()
}
