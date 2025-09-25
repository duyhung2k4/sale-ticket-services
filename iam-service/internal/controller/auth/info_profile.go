package authcontroller

import "github.com/gin-gonic/gin"

func (c *authController) InfoProfile(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(401, map[string]any{
			"err": "token not found",
		})
		return
	}

	claims, err := c.jwt.ValidateToken(token)
	if err != nil {
		ctx.JSON(401, map[string]any{
			"err": err.Error(),
		})
		return
	}

	ctx.JSON(200, claims)
}
