package controller

import (
	"net/http"
	"strconv"
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/service"
	"github.com/gin-gonic/gin"
)

type TicketController struct {
	service service.TicketService
}

func NewTicketController(service service.TicketService) *TicketController {
	return &TicketController{service}
}

func (c *TicketController) BuyTicket(ctx *gin.Context) {
	var input dto.TicketCreateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetUint("userID")
	ticket, err := c.service.BuyTicket(userID, input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"ticket": ticket})
}

func (c *TicketController) GetTickets(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	tickets, totalPages, totalItems, err := c.service.GetTicketsByUser(userID, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": tickets,
		"pagination": gin.H{
			"current_page": page,
			"total_pages": totalPages,
			"total_items": totalItems,
		},
	})
}

func (c *TicketController) GetTicketDetail(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	ticketID, _ := strconv.Atoi(ctx.Param("id"))
	ticket, err := c.service.GetTicketByID(userID, uint(ticketID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ticket": ticket})
}

func (c *TicketController) CancelTicket(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	ticketID, _ := strconv.Atoi(ctx.Param("id"))
	var input dto.TicketStatusUpdateDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket, err := c.service.UpdateTicketStatus(userID, uint(ticketID), input.Status)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ticket": ticket})
}
