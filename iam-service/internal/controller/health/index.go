package health

import (
	"sale-ticket-iam-service/internal/common/constant"
	"sale-ticket-iam-service/internal/common/middleware"
	"sale-ticket-iam-service/internal/common/utils"

	"github.com/gin-gonic/gin"
)

type healthController struct{}

type HealthController interface {
	Ping(ctx *gin.Context)
}

func NewHandle() HealthController {
	return &healthController{}
}

func ConfigRouter(r *gin.RouterGroup) {
	handle := NewHandle()

	utils.AddRouter(r, utils.AddRouterConfig{
		Method:   constant.GET,
		EndPoint: "/ping",
		Middlewares: []gin.HandlerFunc{
			middleware.Block,
		},
		Controller: handle.Ping,
	})
}
