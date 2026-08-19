package routes

import (
	middleware "github.com/HirunikaGunathunga/GO_REST/Middleware"
	"github.com/gin-gonic/gin"
)

func MainRoutes(server *gin.Engine) {
	server.GET("/Events", getEvents)
	server.GET("/Events/:id", getEventById)
	server.POST("/Events", middleware.Authenticate, createEvents)
	server.PUT("/Events/:id", middleware.Authenticate, updateEventById)
	server.DELETE("/Events/:id", middleware.Authenticate, deleteEventById)

	//users
	server.POST("/SignUp", SignUp)
	server.POST("/Login", Login)
}
