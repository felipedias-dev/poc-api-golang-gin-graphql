package main

import (
	"io"
	"os"

	"github.com/felipedias-dev/poc-api-golang-gin-graphql/controller"
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/middlewares"
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/service"
	"github.com/gin-gonic/gin"
	ginDump "github.com/tpkeeper/gin-dump"
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
	setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		ginDump.Dump(),
	)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello World"})
	})

	server.Run(":8080")
}
