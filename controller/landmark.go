package controller

import (
	"gorm_hex/model"
	"gorm_hex/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLandmarkController(router *gin.Engine) {
	ping := router.Group("/landmark")
	{
		ping.GET("", getAllLandmarks)
		ping.GET(":name", getLandmarksByName)
	}

}

func getAllLandmarks(ctx *gin.Context) {
	searchServ := service.NewSearchService()
	landmarks := []model.Landmark{}
	landmarks, err := searchServ.GetAllLandmarks()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "error wow",
		})
	}
	ctx.JSON(http.StatusOK, landmarks)
}

func getLandmarksByName(ctx *gin.Context) {
	name := ctx.Param("name")
	searchServ := service.NewSearchService()
	landmarks := []model.Landmark{}
	landmarks, err := searchServ.GetLandmarksByName(name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "error wow",
		})
	}
	ctx.JSON(http.StatusOK, landmarks)
}
