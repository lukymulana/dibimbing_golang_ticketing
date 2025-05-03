package service

import (
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/entity"
	"dibimbing_golang_ticketing/repository"
	"errors"
	"time"
)

type TicketService interface {
	BuyTicket(userID uint, dto dto.TicketCreateDTO) (*entity.Ticket, error)
	GetTicketsByUser(userID uint, page, limit int) ([]entity.Ticket, int, int64, error)
	GetTicketByID(userID, ticketID uint) (*entity.Ticket, error)
	UpdateTicketStatus(userID, ticketID uint, status string) (*entity.Ticket, error)
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	eventRepo  repository.EventRepository
}

func NewTicketService(ticketRepo repository.TicketRepository, eventRepo repository.EventRepository) TicketService {
	return &ticketService{ticketRepo, eventRepo}
}

func (s *ticketService) BuyTicket(userID uint, input dto.TicketCreateDTO) (*entity.Ticket, error) {
	event, err := s.eventRepo.GetEventByID(input.EventID)
	if err != nil {
		return nil, errors.New("event not found")
	}
	if event.Status != "Aktif" {
		return nil, errors.New("event is not open for ticket purchase")
	}
	if event.Capacity <= 0 {
		return nil, errors.New("event is full")
	}
	countSold, _ := s.ticketRepo.CountTicketsByEventStatus(event.ID, "tersedia")
	if int(countSold) >= event.Capacity {
		return nil, errors.New("event capacity exceeded")
	}
	ticket := &entity.Ticket{
		EventID:   event.ID,
		UserID:    userID,
		Status:    "tersedia",
		Price:     event.Price,
		CreatedAt: time.Now(),
	}
	if err := s.ticketRepo.CreateTicket(ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *ticketService) GetTicketsByUser(userID uint, page, limit int) ([]entity.Ticket, int, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	tickets, total, err := s.ticketRepo.GetTicketsByUser(userID, page, limit)
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	return tickets, totalPages, total, err
}

func (s *ticketService) GetTicketByID(userID, ticketID uint) (*entity.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return nil, errors.New("ticket not found")
	}
	if ticket.UserID != userID {
		return nil, errors.New("unauthorized access to ticket")
	}
	return ticket, nil
}

func (s *ticketService) UpdateTicketStatus(userID, ticketID uint, status string) (*entity.Ticket, error) {
	ticket, err := s.ticketRepo.GetTicketByID(ticketID)
	if err != nil {
		return nil, errors.New("ticket not found")
	}
	if ticket.UserID != userID {
		return nil, errors.New("unauthorized")
	}
	if ticket.Status == "dibatalkan" {
		return nil, errors.New("ticket already cancelled")
	}
	ticket.Status = status
	if err := s.ticketRepo.UpdateTicket(ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}
