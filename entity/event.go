package entity

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Description string    `json:"description"`
	City        string    `gorm:"not null" json:"city"`
	Capacity    int       `gorm:"not null" json:"capacity"`
	Price       float64   `gorm:"not null" json:"price"`
	Status      string    `gorm:"not null" json:"status"` //enum
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tickets     []Ticket  `gorm:"foreignKey:EventID" json:"tickets"`
}
