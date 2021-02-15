package controllers

import (
	"net/http"
	"time"
	"whatw/database"
	"whatw/models"

	"github.com/gin-gonic/gin"
)

type JsonRequest struct {
	Model string  `json:"model"`
	Maker string  `json:"maker"`
	Size  float32 `json:"size"`
	Low   uint8   `json:"low"`
	Hi    uint8   `json:"hi"`
}
type JsonResponse struct {
	ID uint
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
		"data":    result,
	})
}

func (hc *DisplayController) Put(c *gin.Context) {

	json := &JsonRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	defer db.Close()

	model := models.Display{Maker: json.Maker, Model: json.Model, Size: json.Size, Hi: json.Hi, Low: json.Low,
		CreatedAt: time.Now(), UpdatedAt: time.Now()}
	result := db.Create(&model)

	//c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": http.StatusText(http.StatusOK),
		"result":  "OK",
		"row":     result.RowsAffected,
		"data":    model,
	})
}
