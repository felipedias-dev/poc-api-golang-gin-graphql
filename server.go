package main

import (
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/controller"
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

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
