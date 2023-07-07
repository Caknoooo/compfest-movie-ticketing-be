package entities

import "github.com/google/uuid"

type Ticket struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Amount int       `gorm:"type:int" json:"amount"`
	Price  float64       `gorm:"type:int" json:"price"`

	UserID       uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"-"`
	TicketStok []TicketStok `gorm:"foreignKey:TicketID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	Timestamp
}
