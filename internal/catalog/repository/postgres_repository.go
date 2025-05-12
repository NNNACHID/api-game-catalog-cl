package repository

import (
	"context"
	"errors"
	"math"
	"strings"

	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/models"
	"gorm.io/gorm"
)

type PostgresGameRepository struct {
	db *gorm.DB
}

func NewPostgresGameRepository(db *gorm.DB) GameRepository {
	return &PostgresGameRepository{
		db: db,
	}
}

func (r *PostgresGameRepository) Create(ctx context.Context, game *models.Game) error {
	return r.db.WithContext(ctx).Create(game).Error
}

func (r *PostgresGameRepository) GetByID(ctx context.Context, id uint) (*models.Game, error) {
	var game models.Game
	result := r.db.WithContext(ctx).Preload("Genres").Preload("Platforms").First(&game, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("game not found")
		}
		return nil, result.Error
	}
	return &game, nil
}

func (r *PostgresGameRepository) Update(ctx context.Context, game *models.Game) error {
	return r.db.WithContext(ctx).Save(game).Error
}

func (r *PostgresGameRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Game{}, id).Error
}

func (r *PostgresGameRepository) List(ctx context.Context, filter *models.GameFilter) (*models.GameResponse, error) {
	var games []models.Game
	var totalCount int64
	
	if filter.Page <= 0 {
		filter.Page = 1
	}
	
	if filter.PageSize <= 0 {
		filter.PageSize = 10
	}
	
	query := r.db.WithContext(ctx).Model(&models.Game{}).Preload("Genres").Preload("Platforms")
	
	if filter.Title != "" {
		query = query.Where("title ILIKE ?", "%"+filter.Title+"%")
	}
	if filter.Developer != "" {
		query = query.Where("developer ILIKE ?", "%"+filter.Developer+"%")
	}
	if filter.Publisher != "" {
		query = query.Where("publisher ILIKE ?", "%"+filter.Publisher+"%")
	}
	if filter.MinRating != nil {
		query = query.Where("average_rating >= ?", *filter.MinRating)
	}
	
	if len(filter.Genres) > 0 {
		query = query.Joins("JOIN game_genres ON games.id = game_genres.game_id").
			Joins("JOIN genres ON genres.id = game_genres.genre_id").
			Where("genres.name IN ?", filter.Genres).
			Group("games.id")
	}
	
	if len(filter.Platforms) > 0 {
		query = query.Joins("JOIN game_platforms ON games.id = game_platforms.game_id").
			Joins("JOIN platforms ON platforms.id = game_platforms.platform_id").
			Where("platforms.name IN ?", filter.Platforms).
			Group("games.id")
	}
	
	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	
	if filter.SortBy != "" {
		validFields := map[string]bool{
			"title": true, "release_date": true, "price": true, "average_rating": true, "created_at": true,
		}
		
		sortField := "id"
		if validFields[strings.ToLower(filter.SortBy)] {
			sortField = strings.ToLower(filter.SortBy)
		}
		
		sortOrder := "ASC"
		if strings.ToUpper(filter.SortOrder) == "DESC" {
			sortOrder = "DESC"
		}
		
		query = query.Order(sortField + " " + sortOrder)
	} else {
		query = query.Order("id ASC")
	}
	
	offset := (filter.Page - 1) * filter.PageSize
	err = query.Offset(offset).Limit(filter.PageSize).Find(&games).Error
	if err != nil {
		return nil, err
	}
	
	totalPages := int(math.Ceil(float64(totalCount) / float64(filter.PageSize)))
	
	response := &models.GameResponse{
		Games:      games,
		TotalCount: totalCount,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: totalPages,
	}
	
	return response, nil
}

func (r *PostgresGameRepository) CreateGenre(ctx context.Context, genre *models.Genre) error {
	return r.db.WithContext(ctx).Create(genre).Error
}

func (r *PostgresGameRepository) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	var genres []models.Genre
	err := r.db.WithContext(ctx).Find(&genres).Error
	return genres, err
}

func (r *PostgresGameRepository) CreatePlatform(ctx context.Context, platform *models.Platform) error {
	return r.db.WithContext(ctx).Create(platform).Error
}

func (r *PostgresGameRepository) GetAllPlatforms(ctx context.Context) ([]models.Platform, error) {
	var platforms []models.Platform
	err := r.db.WithContext(ctx).Find(&platforms).Error
	return platforms, err
}
