package http

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", Health)
	r.GET("/greet", Greet)
}
