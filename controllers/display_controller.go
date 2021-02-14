package controllers

import (
	"net/http"
	"whatw/database"
	"whatw/models"

	"github.com/gin-gonic/gin"
)

type JsonRequest struct {
	Model  string `json:"field_model"`
	Size  uint8    `json:"field_size"`
	Low  uint8    `json:"field_low"`
	Hi  uint8    `json:"field_hi"`
}
type JsonResponse struct {
	ID uint8
	JsonRequest
}

// DisplayController controller for health request
type DisplayController struct{}

// NewDisplayController is constructor for DisplayController
func NewDisplayController() *DisplayController {
	database.Init(true, models.Display{})
	return new(DisplayController)
}

// Index is index route for health
func (hc *DisplayController) Index(c *gin.Context) {
	db := database.GetDB()
	defer db.Close()
	var result JsonResponse

	var allDisplay []models.Display
	db.Find(&allDisplay)

	for _, display := range allDisplay {
		result = JsonResponse{
			ID:          display.ID,
			JsonRequest: JsonRequest{},
		}
	}


	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"result":  "OK",
		"data": result,
	})
}

func (hc *DisplayController) Put(c *gin.Context) {

	var json JsonRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	defer db.Close()

	model := models.Display{Model: json.Model, Size: json.Size, Hi: json.Hi, Low: json.Low}
	db.Create(model)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"result":  "OK",
		"data": model,
	})
}
