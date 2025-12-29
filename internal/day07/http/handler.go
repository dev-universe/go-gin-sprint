package http

import (
	"errors"
	nethttp "net/http"

	"github.com/dev-universe/go-gin-sprint/internal/day05"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(nethttp.StatusOK, gin.H{
		"status": "ok",
	})
}

func Greet(c *gin.Context) {
	names := c.QueryArray("name")

	data, err := day05.BuildGreetings(names)
	if err != nil {
		if errors.Is(err, day05.ErrNoValidNames) {
			c.JSON(nethttp.StatusBadRequest, gin.H{"error": "no valid names"})
			return
		}
		if errors.Is(err, day05.ErrNameTooLong) {
			c.JSON(nethttp.StatusUnprocessableEntity, gin.H{
				"error":  "name too long",
				"detail": err.Error(),
			})
			return
		}
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(nethttp.StatusOK, data)
}
