package main

import (
	"errors"
	"net/http"

	"github.com/dev-universe/go-gin-sprint/internal/day05"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/greet", func(c *gin.Context) {
		names := c.QueryArray("name")

		data, err := day05.BuildGreetings(names)

		if err != nil {

			// case : invalid names
			if errors.Is(err, day05.ErrNoValidNames) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "no valid names",
				})
				return
			}

			// case : name too long
			if errors.Is(err, day05.ErrNameTooLong) {
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error":  "name too long",
					"detail": err.Error(),
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error",
			})
			return

		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
