package main

import (
	"io"
	"net/http"
	"os"

	"github.com/felipedias-dev/poc-api-golang-gin-graphql/controller"
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/middlewares"
	"github.com/felipedias-dev/poc-api-golang-gin-graphql/service"
	"github.com/gin-gonic/gin"
	// ginDump "github.com/tpkeeper/gin-dump"
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

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		// ginDump.Dump(),
	)

	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hello World"})
		})

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
