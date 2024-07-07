package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController() HealthController {
	c := HealthController{}
	return c
}

func (c *HealthController) HandleHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Ok")
}
