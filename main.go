package main

import (
	"dibimbing_golang_ticketing/config"
	"dibimbing_golang_ticketing/controller"
	"dibimbing_golang_ticketing/repository"
	"dibimbing_golang_ticketing/service"
	"dibimbing_golang_ticketing/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load .env
	_ = godotenv.Load()

	// Init DB
	db := config.InitDB()

	// Init Repository
	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)
	ticketRepo := repository.NewTicketRepository(db)

	// Init Service
	authService := service.NewAuthService(userRepo)
	eventService := service.NewEventService(eventRepo)
	ticketService := service.NewTicketService(ticketRepo, eventRepo)
	reportService := service.NewReportService(eventRepo, ticketRepo)

	// Init Controller
	authCtrl := controller.NewAuthController(authService)
	eventCtrl := controller.NewEventController(eventService)
	ticketCtrl := controller.NewTicketController(ticketService)
	reportCtrl := controller.NewReportController(reportService)

	// Setup Router
	r := routes.SetupRouter(authCtrl, eventCtrl, ticketCtrl, reportCtrl)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on port", port)
	r.Run(":" + port)
}
