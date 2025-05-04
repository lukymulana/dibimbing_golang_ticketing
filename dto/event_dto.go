package dto

import "time"

type EventCreateDTO struct {
	Name        string    `json:"name" binding:"required,min=3"`
	Description string    `json:"description"`
	City        string    `json:"city" binding:"required"`
	Capacity    int       `json:"capacity" binding:"required,gte=0"`
	Price       float64   `json:"price" binding:"required,gte=0"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

type EventUpdateDTO struct {
	Name        string    `json:"name" binding:"omitempty,min=3"`
	Description string    `json:"description"`
	City        string    `json:"city" binding:"omitempty"`
	Capacity    int       `json:"capacity" binding:"omitempty,gte=0"`
	Price       float64   `json:"price" binding:"omitempty,gte=0"`
	Status      string    `json:"status" binding:"omitempty,oneof=Aktif Berlangsung Selesai"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
