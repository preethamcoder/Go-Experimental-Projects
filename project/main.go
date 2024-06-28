package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", getEventsFr)
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, "Ayo this is how we access something")
}
func getEventsFr(context *gin.Context) {
	context.JSON(http.StatusOK, "Home page")
}
