package service

import (
	"context"

	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/models"
)

type GameService interface {
	CreateGame(ctx context.Context, game *models.Game) error
	GetGameByID(ctx context.Context, id uint) (*models.Game, error)
	UpdateGame(ctx context.Context, game *models.Game) error
	DeleteGame(ctx context.Context, id uint) error
	ListGames(ctx context.Context, filter *models.GameFilter) (*models.GameResponse, error)
	
	CreateGenre(ctx context.Context, genre *models.Genre) error
	GetAllGenres(ctx context.Context) ([]models.Genre, error)
	
	CreatePlatform(ctx context.Context, platform *models.Platform) error
	GetAllPlatforms(ctx context.Context) ([]models.Platform, error)
}
