package routes

import (
	"dibimbing_golang_ticketing/controller"
	"dibimbing_golang_ticketing/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authCtrl *controller.AuthController,
	eventCtrl *controller.EventController,
	ticketCtrl *controller.TicketController,
	reportCtrl *controller.ReportController,
) *gin.Engine {
	r := gin.Default()

	// Auth
	r.POST("/register", authCtrl.Register)
	r.POST("/login", authCtrl.Login)

	// Event
	event := r.Group("/events")
	event.GET("", eventCtrl.GetEvents)
	event.Use(middleware.AuthMiddleware(), middleware.RequireRole("admin"))
	event.POST("", eventCtrl.CreateEvent)
	event.PUT(":id", eventCtrl.UpdateEvent)
	event.DELETE(":id", eventCtrl.DeleteEvent)

	// Ticket
	ticket := r.Group("/tickets", middleware.AuthMiddleware(), middleware.RequireRole("user"))
	ticket.GET("", ticketCtrl.GetTickets)
	ticket.POST("", ticketCtrl.BuyTicket)
	ticket.GET(":id", ticketCtrl.GetTicketDetail)
	ticket.PATCH(":id", ticketCtrl.CancelTicket)

	// Report
	report := r.Group("/reports", middleware.AuthMiddleware(), middleware.RequireRole("admin"))
	report.GET("/summary", reportCtrl.GetSummary)
	report.GET("/event/:id", reportCtrl.GetEventReport)

	return r
}
