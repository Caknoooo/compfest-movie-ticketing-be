package entities

import (
	"time"

	"github.com/google/uuid"
)

type Topup struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Jumlah           float64   `gorm:"type:float" json:"jumlah"`
	TanggalTransaksi time.Time `gorm:"timestamp with time zone" json:"tanggal_transaksi"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID" json:"-"`
}
