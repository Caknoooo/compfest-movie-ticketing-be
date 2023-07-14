package dto

type CreateWithDrawalDTO struct {
	JumlahPenarikan float64 `gorm:"type:float" json:"jumlah_penarikan"`
	BankID          uint    `gorm:"type:uint" json:"bank_id"`
}