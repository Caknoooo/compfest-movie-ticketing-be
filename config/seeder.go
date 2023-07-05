package config

import (
	"errors"

	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListBankSeeder(db); err != nil {
		return err
	}

	return nil
}

func ListBankSeeder(db *gorm.DB) error {
	var listBank = []entities.ListBank{
		{
			ID:   1,
			Name: "BCA",
		}, 
		{
			ID:   2,
			Name: "BNI",
		}, 
		{
			ID:   3,
			Name: "BRI",
		}, 
		{
			ID:   4,
			Name: "Mandiri",
		},
		{
			ID:   5,
			Name: "OVO",
		}, 
		{
			ID:   6,
			Name: "Gopay",
		}, 
	}

	hasTable := db.Migrator().HasTable(&entities.ListBank{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.ListBank{}); err != nil {
			return err
		}
	}

	for _, data := range listBank {
		var bank entities.ListBank
		err := db.Where(&entities.ListBank{ID: data.ID}).First(&bank).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}