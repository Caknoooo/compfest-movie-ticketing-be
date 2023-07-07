package dto

type TicketCreateDTO struct {
	Amount int `json:"amount" binding:"required"`
	Price  float64 `json:"price" binding:"required"`

	MovieID string `json:"movie_id" binding:"required"`
}