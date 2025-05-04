package repository

import (
	"dibimbing_golang_ticketing/entity"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *entity.Event) error
	GetAllEvents(keyword, status, city, startDate, endDate string, page, limit int) ([]entity.Event, int64, error)
	GetEventByID(id uint) (*entity.Event, error)
	UpdateEvent(event *entity.Event) error
	DeleteEvent(id uint) error
	IsEventNameUnique(name string, excludeID uint) (bool, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) CreateEvent(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) GetAllEvents(keyword, status, city string, startDate, endDate string, page, limit int) ([]entity.Event, int64, error) {
	var events []entity.Event
	var count int64
	q := r.db.Model(&entity.Event{})
	if keyword != "" {
		q = q.Where("name LIKE ?", "%"+keyword+"%")
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if city != "" {
		q = q.Where("city = ?", city)
	}
	if startDate != "" && endDate != "" {
		q = q.Where("start_date >= ? AND end_date <= ?", startDate, endDate)
	} else if startDate != "" {
		q = q.Where("start_date >= ?", startDate)
	} else if endDate != "" {
		q = q.Where("end_date <= ?", endDate)
	}
	q.Count(&count)
	offset := (page - 1) * limit
	if err := q.Order("start_date desc").Limit(limit).Offset(offset).Find(&events).Error; err != nil {
		return nil, 0, err
	}
	return events, count, nil
}

func (r *eventRepository) GetEventByID(id uint) (*entity.Event, error) {
	var event entity.Event
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) UpdateEvent(event *entity.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&entity.Event{}, id).Error
}

func (r *eventRepository) IsEventNameUnique(name string, excludeID uint) (bool, error) {
	var count int64
	q := r.db.Model(&entity.Event{}).Where("name = ?", name)
	if excludeID != 0 {
		q = q.Where("id <> ?", excludeID)
	}
	if err := q.Count(&count).Error; err != nil {
		return false, err
	}
	return count == 0, nil
}
