package router

import (
	"sale-ticket-iam-service/internal/controller/health"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func handleApiV1(r *gin.RouterGroup) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	health.ConfigRouter(r)
}
