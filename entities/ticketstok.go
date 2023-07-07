package entities

import "github.com/google/uuid"

type TicketStok struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Nomor uint `gorm:"type:int" json:"nomor"`

	TicketID uuid.UUID `gorm:"type:uuid" json:"ticket_id"`
	Ticket   Ticket  `gorm:"foreignKey:TicketID" json:"-"`
	MovieID uuid.UUID `gorm:"type:uuid" json:"movie_id"`
	Movie   Movies  `gorm:"foreignKey:MovieID" json:"-"`
}