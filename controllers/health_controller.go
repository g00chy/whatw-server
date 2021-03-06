package controllers

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController controller for health request
type HealthController struct{}

// NewHealthController is constructor for HealthController
func NewHealthController() *HealthController {
	return new(HealthController)
}

// Index is index route for health
func (hc *HealthController) Index(c echo.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"result":  "OK",
	})
	return nil
}
