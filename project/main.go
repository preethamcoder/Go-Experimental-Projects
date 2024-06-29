package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", getHome)
	server.POST("/events", createEvent)

	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}
func getHome(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "required values missing"})
	}
}
