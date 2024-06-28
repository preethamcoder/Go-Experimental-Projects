package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", getHome)
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, "Ayo this is how we access something")
}
func getHome(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
