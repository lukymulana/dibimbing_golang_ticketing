package dto

type ReportSummaryDTO struct {
	TotalTickets int     `json:"total_tickets"`
	TotalRevenue float64 `json:"total_revenue"`
}

type ReportEventDTO struct {
	EventID      uint    `json:"event_id"`
	EventName    string  `json:"event_name"`
	TicketsSold  int     `json:"tickets_sold"`
	Revenue      float64 `json:"revenue"`
}
