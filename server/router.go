package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"whatw/config"
	"whatw/controllers"
)

// NewRouter is constructor for router
func NewRouter() (*gin.Engine, error) {
	c := config.GetConfig()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	//version := router.Group("/" + c.GetString("server.version"))

	healthController := controllers.NewHealthController()
	router.GET("/health", healthController.Index)

	return router, nil
}
