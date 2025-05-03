package controller

import (
	"net/http"
	"strconv"
	"time"
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/service"
	"github.com/gin-gonic/gin"
)

type EventController struct {
	service service.EventService
}

func NewEventController(service service.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var input dto.EventCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := c.service.CreateEvent(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"event": event})
}

func (c *EventController) GetEvents(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	status := ctx.Query("status")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	events, totalPages, totalItems, err := c.service.GetAllEvents(keyword, status, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": events,
		"pagination": gin.H{
			"current_page": page,
			"total_pages": totalPages,
			"total_items": totalItems,
		},
	})
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.EventUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := c.service.UpdateEvent(uint(id), input, time.Now())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"event": event})
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.service.DeleteEvent(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
