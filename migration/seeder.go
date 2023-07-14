package migration

import (
	"errors"
	"fmt"
	"time"

	"github.com/Caknoooo/golang-clean_template/dto"
	"github.com/Caknoooo/golang-clean_template/entities"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := ListBankSeeder(db); err != nil {
		return err
	}

	if err := ListUserSeeder(db); err != nil {
		return err
	}

	if err := ListMovieSeeder(db); err != nil {
		return err
	}

	if err := ListTimeMovieSeeder(db); err != nil {
		return err
	}

	if err := ListStudioSeeder(db); err != nil {
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

func ListTimeMovieSeeder(db *gorm.DB) error {
	var listTimeMovie = []entities.TimeMovie{
		{
			ID:   1,
			Time: "12.05",
			Type: 1,
		},
		{
			ID:   2,
			Time: "14.30",
			Type: 1,
		},
		{
			ID:   3,
			Time: "16.05",
			Type: 1,
		},
		{
			ID:   4,
			Time: "18.30",
			Type: 1,
		},
		{
			ID:   5,
			Time: "20.05",
			Type: 1,
		},
		{
			ID:   6,
			Time: "22.30",
			Type: 1,
		},
		{
			ID:   7,
			Time: "10.05",
			Type: 2,
		},
		{
			ID:   8,
			Time: "12.30",
			Type: 2,
		},
		{
			ID:   9,
			Time: "14.05",
			Type: 2,
		},
		{
			ID:   10,
			Time: "16.30",
			Type: 2,
		},
		{
			ID:   11,
			Time: "18.05",
			Type: 2,
		},
		{
			ID:   12,
			Time: "20.30",
			Type: 2,
		},
		{
			ID:   13,
			Time: "22.05",
			Type: 2,
		},
		{
			ID:   14,
			Time: "09.05",
			Type: 3,
		},
		{
			ID:   15,
			Time: "11.30",
			Type: 3,
		},
		{
			ID:   16,
			Time: "13.05",
			Type: 3,
		},
		{
			ID:   17,
			Time: "15.30",
			Type: 3,
		},
		{
			ID:   18,
			Time: "17.05",
			Type: 3,
		},
		{
			ID:   19,
			Time: "19.30",
			Type: 3,
		},
		{
			ID:   20,
			Time: "11.05",
			Type: 4,
		},
		{
			ID:   21,
			Time: "13.30",
			Type: 4,
		},
		{
			ID:   22,
			Time: "15.05",
			Type: 4,
		},
		{
			ID:   23,
			Time: "17.30",
			Type: 4,
		},
		{
			ID:   24,
			Time: "19.05",
			Type: 4,
		},
		{
			ID:   25,
			Time: "21.30",
			Type: 4,
		},
		{
			ID:   26,
			Time: "10.05",
			Type: 5,
		},
		{
			ID:   27,
			Time: "12.30",
			Type: 5,
		},
		{
			ID:   28,
			Time: "14.05",
			Type: 5,
		},
		{
			ID:   29,
			Time: "16.30",
			Type: 5,
		},
		{
			ID:   30,
			Time: "18.05",
			Type: 5,
		},
	}

	for _, data := range listTimeMovie {
		var timeMovie entities.TimeMovie
		err := db.Where(&entities.TimeMovie{ID: data.ID}).First(&timeMovie).Error
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

func ListStudioSeeder(db *gorm.DB) error {
	var listPlace = []entities.Place{
		{
			ID:   1,
			Name: "CGV",
		},
		{
			ID:   2,
			Name: "XXI",
		},
		{
			ID:   3,
			Name: "Cinepolis",
		},
		{
			ID:   4,
			Name: "Cinemaxx",
		},
		{
			ID:   5,
			Name: "New Star Cineplex",
		},
		{
			ID:   6,
			Name: "Flix Cinema",
		},
		{
			ID:   7,
			Name: "Hiflix",
		},
		{
			ID:   8,
			Name: "Platinum Cineplex",
		},
	}

	for _, data := range listPlace {
		var place entities.Place
		err := db.Where(&entities.Place{ID: data.ID}).First(&place).Error
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

func ListUserSeeder(db *gorm.DB) error {
	var listUser = []entities.User{
		{
			Nama:         "Admin",
			NoTelp:       "081234567890",
			Email:        "Admin@gmail.com",
			Password:     "Admin123",
			Age:          20,
			TanggalLahir: "08/15/2003",
			Role:         "Admin",
			Saldo:        0,
		},
		{
			Nama:         "User",
			NoTelp:       "081234567890",
			Email:        "User@gmail.com",
			Password:     "User123",
			Age:          20,
			TanggalLahir: "08/15/2003",
			Role:         "User",
			Saldo:        0,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.User{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.User{}); err != nil {
			return err
		}
	}

	for _, data := range listUser {
		var user entities.User
		err := db.Where(&entities.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func ListMovieSeeder(db *gorm.DB) error {
	var listMovie = []dto.MovieCreateDTO{
		{
			Title:       "Fast X",
			Description: "Dom Toretto dan keluarganya menjadi sasaran putra gembong narkoba Hernan Reyes yang pendendam.",
			ReleaseDate: "2023-05-17",
			PosterUrl:   "https://image.tmdb.org/t/p/w500/fiVW06jE7z9YnO4trhaMEdclSiC.jpg",
			AgeRating:   15,
			TicketPrice: 53000,
		},
	}

	hasTable := db.Migrator().HasTable(&entities.Movies{})
	if !hasTable {
		if err := db.AutoMigrate(&entities.Movies{}); err != nil {
			return err
		}
	}

	for _, data := range listMovie {
		var movie entities.Movies
		err := db.Where(&entities.Movies{Title: data.Title}).First(&movie).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			t, err := time.Parse("2006-01-02", data.ReleaseDate)
			if err != nil {
				fmt.Println("Error parsing time:", err)
				return err
			}

			timeData := t.In(time.FixedZone("WIB", 7*60*60))

			var dataMovie = entities.Movies{
				Title:       data.Title,
				Description: data.Description,
				ReleaseDate: timeData,
				PosterUrl:   data.PosterUrl,
				AgeRating:   data.AgeRating,
				TicketPrice: data.TicketPrice,
			}

			if err := db.Create(&dataMovie).Error; err != nil {
				return err
			}
		} else {
			fmt.Println("Movie already exists:", movie.Title)
		}
	}

	return nil
}
