package healthcontroller

import (
	"sale-ticket-iam-service/internal/common/constant"
	"sale-ticket-iam-service/internal/common/utils"

	"github.com/gin-gonic/gin"
)

type healthController struct{}

type HealthController interface {
	Ping(ctx *gin.Context)
}

func MakeHandle() HealthController {
	return &healthController{}
}

func Register(r *gin.RouterGroup) {
	handle := MakeHandle()

	utils.AddRouter(r, utils.AddRouterConfig{
		Method:      constant.GET,
		EndPoint:    "/ping",
		Middlewares: []gin.HandlerFunc{},
		Controller:  handle.Ping,
	})
}
