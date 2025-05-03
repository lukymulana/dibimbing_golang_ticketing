package service

import (
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/repository"
)

type ReportService interface {
	GetSummaryReport() (*dto.ReportSummaryDTO, error)
	GetEventReport(eventID uint) (*dto.ReportEventDTO, error)
}

type reportService struct {
	eventRepo  repository.EventRepository
	ticketRepo repository.TicketRepository
}

func NewReportService(eventRepo repository.EventRepository, ticketRepo repository.TicketRepository) ReportService {
	return &reportService{eventRepo, ticketRepo}
}

func (s *reportService) GetSummaryReport() (*dto.ReportSummaryDTO, error) {
	// Hitung total tiket terjual dan total pendapatan
	var totalTickets int64
	var totalRevenue float64
	events, _, err := s.eventRepo.GetAllEvents("", "", 1, 10000)
	if err != nil {
		return nil, err
	}
	for _, ev := range events {
		count, _ := s.ticketRepo.CountTicketsByEventStatus(ev.ID, "tersedia")
		totalTickets += count
		totalRevenue += float64(count) * ev.Price
	}
	return &dto.ReportSummaryDTO{
		TotalTickets: int(totalTickets),
		TotalRevenue: totalRevenue,
	}, nil
}

func (s *reportService) GetEventReport(eventID uint) (*dto.ReportEventDTO, error) {
	event, err := s.eventRepo.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}
	count, _ := s.ticketRepo.CountTicketsByEventStatus(eventID, "tersedia")
	revenue := float64(count) * event.Price
	return &dto.ReportEventDTO{
		EventID:     event.ID,
		EventName:   event.Name,
		TicketsSold: int(count),
		Revenue:     revenue,
	}, nil
}
