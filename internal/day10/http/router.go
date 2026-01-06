package http

import (
	"errors"
	nethttp "net/http"

	"github.com/dev-universe/go-gin-sprint/internal/day05"
	"github.com/gin-gonic/gin"
)

type Greeter interface {
	BuildGreetings(names []string) ([]day05.Greeting, error)
}

type Handler struct {
	greeter Greeter
}

func NewHandler(g Greeter) *Handler {
	return &Handler{greeter: g}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(nethttp.StatusOK, gin.H{"status": "ok"})
}

func (h Handler) Greet(c *gin.Context) {
	names := c.QueryArray("name")

	data, err := h.greeter.BuildGreetings(names)
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
