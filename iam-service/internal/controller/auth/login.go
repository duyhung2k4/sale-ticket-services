package authcontroller

import (
	"sale-ticket-iam-service/internal/common/utils"

	"github.com/gin-gonic/gin"
)

func (c *authController) Login(ctx *gin.Context) {
	data := utils.ClaimJWT{
		Uuid: "0001",
		Role: "admin",
	}
	token, err := c.jwt.CreateToken(data)
	if err != nil {
		ctx.JSON(502, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, map[string]any{
		"token": token,
	})
}
