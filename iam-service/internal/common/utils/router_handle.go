package utils

import (
	"sale-ticket-iam-service/internal/common/constant"

	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup, config AddRouterConfig) {
	handleFuncs := []gin.HandlerFunc{}
	handleFuncs = append(handleFuncs, config.Middlewares...)
	handleFuncs = append(handleFuncs, config.Controller)
	r.Handle(string(config.Method), config.EndPoint, handleFuncs...)
}

type AddRouterConfig struct {
	Method      constant.METHOD
	EndPoint    string
	Middlewares []gin.HandlerFunc
	Controller  gin.HandlerFunc
}
