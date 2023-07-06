package repository

import (
	"context"

	"github.com/Caknoooo/golang-clean_template/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateMovies(ctx context.Context, movies entities.Movies) (entities.Movies, error)
	GetAllMovies(ctx context.Context) ([]entities.Movies, error)
	GetMovieByID(ctx context.Context, movieID uuid.UUID) (entities.Movies, error)
}

type movieRepository struct {
	connection *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{
		connection: db,
	}
}

func (mr *movieRepository) CreateMovies(ctx context.Context, movies entities.Movies) (entities.Movies, error) {
	if err := mr.connection.Create(&movies).Error; err != nil {
		return entities.Movies{}, err
	}
	return movies, nil
}

func (mr *movieRepository) GetAllMovies(ctx context.Context) ([]entities.Movies, error) {
	var movies []entities.Movies
	if err := mr.connection.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (mr *movieRepository) GetMovieByID(ctx context.Context, movieID uuid.UUID) (entities.Movies, error) {
	var movies entities.Movies
	if err := mr.connection.Where("id = ?", movieID).Take(&movies).Error; err != nil {
		return entities.Movies{}, err
	}
	return movies, nil
}