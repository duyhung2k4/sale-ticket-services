package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppRouter() http.Handler {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	handleApiV1(apiV1)

	return r
}
