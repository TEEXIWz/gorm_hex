package controller

import (
	"gorm_hex/controller/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDemoController(router *gin.Engine) {
	ping := router.Group("/ping")
	{
		ping.GET("", demoPing)
		ping.GET("/:name", demoPingName)
		ping.POST("/postform", demoPostForm)
		ping.POST("/postjson", demoPostJson)
	}

}

func demoPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func demoPingName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
	})
}

func demoPostForm(ctx *gin.Context) {
	name := ctx.Param("name")
	nickname := ctx.DefaultPostForm("nickname", "Unknow")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong " + name + " " + nickname,
	})
}

func demoPostJson(ctx *gin.Context) {
	person := model.Person{}
	ctx.ShouldBindJSON(&person)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong " + person.Name + " " + strconv.Itoa(person.Age),
	})
}
