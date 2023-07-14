package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TopupRepository interface {
	CreateTopup(ctx context.Context, topup entities.Topup) (entities.Topup, error)
	GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error)
	GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error)
}

type topupRepository struct {
	connection *gorm.DB
}

func NewTopupRepository(db *gorm.DB) TopupRepository {
	return &topupRepository{
		connection: db,
	}
}

func (tr *topupRepository) CreateTopup(ctx context.Context, topup entities.Topup) (entities.Topup, error) {
	var user entities.User

	if err := tr.connection.Where("id", topup.UserID).First(&user).Error; err != nil {
		return entities.Topup{}, err
	}

	user.Saldo += topup.Jumlah
	tr.connection.Save(&user)

	if err := tr.connection.Create(&topup).Error; err != nil {
		return entities.Topup{}, err
	}
	return topup, nil
}

func (tr *topupRepository) GetAllTopupUser(ctx context.Context, userID uuid.UUID) ([]entities.Topup, error) {
	var topup []entities.Topup
	if err := tr.connection.Where("user_id", userID).Find(&topup).Error; err != nil {
		return nil, err
	}
	return topup, nil
}

func (tr *topupRepository) GetTopupByID(ctx context.Context, topupID uuid.UUID) (entities.Topup, error) {
	var topup entities.Topup
	if err := tr.connection.Where("id = ?", topupID).Take(&topup).Error; err != nil {
		return entities.Topup{}, err
	}
	return topup, nil
}