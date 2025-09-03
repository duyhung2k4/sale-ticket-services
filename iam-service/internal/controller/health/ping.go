package health

import "github.com/gin-gonic/gin"

func (c *healthController) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]any{
		"mess": "pong",
	})
}
