package authcontroller

import (
	"sale-ticket-iam-service/internal/common/connection"
	"sale-ticket-iam-service/internal/common/utils"

	"github.com/gin-gonic/gin"
)

type authController struct {
	jwt *utils.JWT
}

type AuthController interface {
	Login(ctx *gin.Context)
	InfoProfile(ctx *gin.Context)
}

func MakeHandle() AuthController {
	jwt := &utils.JWT{
		Key: []byte(connection.GetConfig().JwtKey),
	}
	return &authController{
		jwt: jwt,
	}
}

func Register(r *gin.RouterGroup) {
	handle := MakeHandle()

	utils.AddRouter(r, utils.AddRouterConfig{
		Method:      "GET",
		EndPoint:    "/login",
		Middlewares: []gin.HandlerFunc{},
		Controller:  handle.Login,
	})

	utils.AddRouter(r, utils.AddRouterConfig{
		Method:      "GET",
		EndPoint:    "/info-profile",
		Middlewares: []gin.HandlerFunc{},
		Controller:  handle.InfoProfile,
	})
}
