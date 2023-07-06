package dto

type MovieCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	ReleaseDate string `json:"release_date" form:"release_date" binding:"required"`
	AgeRating   int    `json:"age_rating" form:"age_rating" binding:"required"`
	TicketPrice       float64    `json:"ticket_price" form:"price" binding:"required"`
}

type MovieUpdateDTO struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	ReleaseDate string `json:"release_date" form:"release_date"`
	AgeRating   int    `json:"age_rating" form:"age_rating"`
	TicketPrice       float64    `json:"ticket_price" form:"price"`
}