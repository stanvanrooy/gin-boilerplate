package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gin-boilerplate/api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:8001"},
    AllowMethods:     []string{"*"},
    AllowHeaders:     []string{"*"},
    ExposeHeaders:    []string{"*"},
    AllowCredentials: true,
  }))

	router.Use(gin.Recovery())

  router.GET("/ping", controllers.Ping)
	return router
}

