package repository

import (
	"dibimbing_golang_ticketing/entity"
	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ticket *entity.Ticket) error
	GetTicketsByUser(userID uint, page, limit int) ([]entity.Ticket, int64, error)
	GetTicketByID(id uint) (*entity.Ticket, error)
	GetTicketsByEvent(eventID uint) ([]entity.Ticket, error)
	UpdateTicket(ticket *entity.Ticket) error
	CountTicketsByEvent(eventID uint) (int64, error)
	CountTicketsByEventStatus(eventID uint, status string) (int64, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) CreateTicket(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) GetTicketsByUser(userID uint, page, limit int) ([]entity.Ticket, int64, error) {
	var tickets []entity.Ticket
	var count int64
	q := r.db.Model(&entity.Ticket{}).Where("user_id = ?", userID)
	q.Count(&count)
	offset := (page - 1) * limit
	if err := q.Order("created_at desc").Limit(limit).Offset(offset).Find(&tickets).Error; err != nil {
		return nil, 0, err
	}
	return tickets, count, nil
}

func (r *ticketRepository) GetTicketByID(id uint) (*entity.Ticket, error) {
	var ticket entity.Ticket
	if err := r.db.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) GetTicketsByEvent(eventID uint) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	if err := r.db.Where("event_id = ?", eventID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) UpdateTicket(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) CountTicketsByEvent(eventID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Ticket{}).Where("event_id = ?", eventID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ticketRepository) CountTicketsByEventStatus(eventID uint, status string) (int64, error) {
	var count int64
	if err := r.db.Model(&entity.Ticket{}).Where("event_id = ? AND status = ?", eventID, status).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
