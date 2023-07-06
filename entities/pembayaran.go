package entities

import "github.com/google/uuid"

type Pembayaran struct {
	ID     uuid.UUID    `gorm:"type:int;primary_key;auto_increment" json:"id"`
	Jumlah float64 `gorm:"type:float" json:"jumlah"`

	ListBankID uint `gorm:"type:uint" json:"list_bank_id"`
	UserID     uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"user,omitempty"`

	Timestamp
}