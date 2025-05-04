package entity

import "time"

type Ticket struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `gorm:"not null" json:"event_id"`
	Event     Event     `gorm:"foreignKey:EventID" json:"event"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Status    string    `gorm:"not null" json:"status"`
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
