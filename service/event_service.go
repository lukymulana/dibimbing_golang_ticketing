package service

import (
	"dibimbing_golang_ticketing/dto"
	"dibimbing_golang_ticketing/entity"
	"dibimbing_golang_ticketing/repository"
	"errors"
	"time"
)

type EventService interface {
	CreateEvent(dto.EventCreateDTO) (*entity.Event, error)
	GetAllEvents(keyword, status, city, startDate, endDate string, page, limit int) ([]entity.Event, int, int64, error)
	GetEventByID(id uint) (*entity.Event, error)
	UpdateEvent(id uint, dto dto.EventUpdateDTO, currentTime time.Time) (*entity.Event, error)
	DeleteEvent(id uint) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) CreateEvent(input dto.EventCreateDTO) (*entity.Event, error) {
	unique, err := s.repo.IsEventNameUnique(input.Name, 0)
	if err != nil {
		return nil, err
	}
	if !unique {
		return nil, errors.New("event name must be unique")
	}
	if input.Capacity < 0 {
		return nil, errors.New("event capacity cannot be negative")
	}
	if input.Price < 0 {
		return nil, errors.New("event price cannot be negative")
	}
	event := &entity.Event{
		Name:        input.Name,
		Description: input.Description,
		Capacity:    input.Capacity,
		Price:       input.Price,
		Status:      "Aktif",
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.CreateEvent(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *eventService) GetAllEvents(keyword, status, city, startDate, endDate string, page, limit int) ([]entity.Event, int, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	events, total, err := s.repo.GetAllEvents(keyword, status, city, startDate, endDate, page, limit)
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	return events, totalPages, total, err
}

func (s *eventService) GetEventByID(id uint) (*entity.Event, error) {
	return s.repo.GetEventByID(id)
}

func (s *eventService) UpdateEvent(id uint, input dto.EventUpdateDTO, currentTime time.Time) (*entity.Event, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		return nil, errors.New("event not found")
	}
	if event.Status == "Selesai" || currentTime.After(event.EndDate) {
		return nil, errors.New("event already finished, cannot be updated")
	}
	if input.Name != "" {
		unique, _ := s.repo.IsEventNameUnique(input.Name, id)
		if !unique {
			return nil, errors.New("event name must be unique")
		}
		event.Name = input.Name
	}
	if input.City != "" {
		event.City = input.City
	}
	if input.Capacity != 0 {
		event.Capacity = input.Capacity
	}
	if input.Price != 0 {
		event.Price = input.Price
	}
	if input.Status != "" {
		event.Status = input.Status
	}
	if !input.StartDate.IsZero() {
		event.StartDate = input.StartDate
	}
	if !input.EndDate.IsZero() {
		event.EndDate = input.EndDate
	}
	if input.Description != "" {
		event.Description = input.Description
	}
	event.UpdatedAt = currentTime
	if err := s.repo.UpdateEvent(event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *eventService) DeleteEvent(id uint) error {
	return s.repo.DeleteEvent(id)
}
