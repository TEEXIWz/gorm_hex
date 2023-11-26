package main

import (
	"gorm_hex/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	controller.NewDemoController(router)
	controller.NewLandmarkController(router)
	router.Run()
}
