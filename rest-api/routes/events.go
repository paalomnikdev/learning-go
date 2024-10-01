package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch event"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't fetch events"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request data."})
		return
	}

	event.UserID = ctx.GetInt64("userId")

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't save event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't fetch event"})
		return
	}

	if event.UserID != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Only creator can update event"})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request data."})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't update event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated.", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't fetch event"})
		return
	}

	if event.UserID != ctx.GetInt64("userId") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Only creator can delete event"})
		return
	}

	err = event.Delete()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Can't delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted."})
}