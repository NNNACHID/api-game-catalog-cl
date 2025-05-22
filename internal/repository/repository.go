package repository

import (
	"context"

	"github.com/NNNACHID/api-game-catalog-cl/internal/models"
)

type GameRepository interface {
	Create(ctx context.Context, game *models.Game) error
	GetByID(ctx context.Context, id uint) (*models.Game, error)
	Update(ctx context.Context, game *models.Game) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filter *models.GameFilter) (*models.GameResponse, error)
	
	CreateGenre(ctx context.Context, genre *models.Genre) error
	GetAllGenres(ctx context.Context) ([]models.Genre, error)
	
	CreatePlatform(ctx context.Context, platform *models.Platform) error
	GetAllPlatforms(ctx context.Context) ([]models.Platform, error)
}
