package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"whatw/database"
	"whatw/models"

	"github.com/gin-gonic/gin"
)

type DisplayJsonRequest struct {
	Model string  `json:"model"`
	Maker string  `json:"maker"`
	Size  float32 `json:"size"`
	Low   uint8   `json:"low"`
	Hi    uint8   `json:"hi"`
}
type JsonResponse struct {
	ID uint
	DisplayJsonRequest
}

// DisplayController controller for health request
type DisplayController struct{}

// NewDisplayController is constructor for DisplayController
func NewDisplayController() *DisplayController {
	database.Init()
	return new(DisplayController)
}

// Index is index route for health
func (hc *DisplayController) Index(c *gin.Context) {
	db := database.GetDB()

	var result JsonResponse

	var allDisplay []models.Display
	db.Find(&allDisplay)

	for _, display := range allDisplay {
		result = JsonResponse{
			ID:                 display.ID,
			DisplayJsonRequest: DisplayJsonRequest{},
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

	buf := make([]byte, 2048)
	// エラー判定がうまく行かないのでスルーします
	length, _ := c.Request.Body.Read(buf)

	defer c.Request.Body.Close()

	if doCheckJson(buf[:length]) != nil {

	}
	var request DisplayJsonRequest
	if err := json.Unmarshal(buf[:length], &request); err != nil {
		log.Fatal(err)
	}

	// TODO: move service
	db := database.GetDB()

	model := models.Display{Maker: request.Maker, Model: request.Model, Size: request.Size, Hi: request.Hi, Low: request.Low,
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

func doCheckJson([]byte) error {

	return nil
}
