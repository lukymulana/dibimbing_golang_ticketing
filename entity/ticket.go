package entity

import "time"

type Ticket struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `gorm:"not null" json:"event_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Status    string    `gorm:"not null" json:"status"` // tersedia, habis, dibatalkan
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
