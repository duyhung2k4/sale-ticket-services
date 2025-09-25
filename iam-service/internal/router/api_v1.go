package router

import (
	authcontroller "sale-ticket-iam-service/internal/controller/auth"
	healthcontroller "sale-ticket-iam-service/internal/controller/health"
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

	healthGroup := r.Group("/health")
	healthcontroller.Register(healthGroup)

	authGroup := r.Group("/auth")
	authcontroller.Register(authGroup)
}
