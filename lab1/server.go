package main

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"lab1/controller"
	"lab1/middleware"
	"lab1/service"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput() // Логи и консоль и в файл gin.log

	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())

	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!!"})
		}

	})

	server.Run(":8080")

}
