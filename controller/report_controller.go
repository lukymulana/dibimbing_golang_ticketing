package controller

import (
	"net/http"
	"strconv"
	"dibimbing_golang_ticketing/service"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service service.ReportService
}

func NewReportController(service service.ReportService) *ReportController {
	return &ReportController{service}
}

func (c *ReportController) GetSummary(ctx *gin.Context) {
	report, err := c.service.GetSummaryReport()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"report": report})
}

func (c *ReportController) GetEventReport(ctx *gin.Context) {
	eventID, _ := strconv.Atoi(ctx.Param("id"))
	report, err := c.service.GetEventReport(uint(eventID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"report": report})
}
