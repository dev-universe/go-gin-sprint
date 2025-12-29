package main

import (
	"github.com/dev-universe/go-gin-sprint/internal/day07/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	http.RegisterRoutes(r)
	r.Run(":8080")
}
