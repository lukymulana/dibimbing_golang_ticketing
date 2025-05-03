package dto

type TicketCreateDTO struct {
	EventID uint `json:"event_id" binding:"required"`
}

type TicketStatusUpdateDTO struct {
	Status string `json:"status" binding:"required,oneof=cancelled"`
}
