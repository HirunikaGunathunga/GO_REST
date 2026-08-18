package routes

import "github.com/gin-gonic/gin"

func MainRoutes(server *gin.Engine) {
	server.GET("/Events", getEvents)
	server.GET("/Events/:id", getEventById)
	server.POST("/Events", createEvents)
	server.PUT("/Events/:id", updateEventById)
}
