package routes

import (
	"fmt"
	"net/http"
	"strconv"

	model "github.com/HirunikaGunathunga/GO_REST/Model"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Cannot fetch. Try again"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Cannot not parse event id. Try again"})
		return
	}
	e, err := model.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Cannot not fetch event id. Try again"})
		return
	}
	context.JSON(http.StatusOK, e)
}

func createEvents(context *gin.Context) {
	var e model.Event
	err := context.ShouldBindJSON(&e)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the action"})
		return
	}
	//e.ID = 1
	//e.UserId = 1
	err = e.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event: ": e})
}

func updateEventById(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Cannot not parse event id. Try again"})
		return
	}
	_, err = model.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Cannot not fetch event id. Try again"})
		return
	}
	var updateEvent model.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Cannot not parse requested id. Try again"})
		return
	}
	updateEvent.ID = eventId
	updateEvent.ID = eventId

	// Check for errors returned by update
	err = updateEvent.UpdateEventById()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message:": "Updated"})
}
