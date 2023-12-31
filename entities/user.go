package entities

import (
	"github.com/Caknoooo/golang-clean_template/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama         string    `gorm:"type:varchar(100)" json:"nama"`
	NoTelp       string    `gorm:"type:varchar(30)" json:"no_telp"`
	Email        string    `gorm:"type:varchar(100)" json:"email"`
	Password     string    `gorm:"type:varchar(100)" json:"password"`
	Age          int       `gorm:"type:int" json:"age"`
	TanggalLahir string    `gorm:"type:varchar(100)" json:"tanggal_lahir"`
	Role         string    `gorm:"type:varchar(100)" json:"role"`
	Saldo        float64   `gorm:"type:float" json:"saldo"`
	Timestamp

	Topup      []Topup      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Pembayaran []Pembayaran `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Ticket     []Ticket     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
