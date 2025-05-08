package service

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/models"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/repository"
)

type gameService struct {
	repo   repository.GameRepository
	logger *logrus.Logger
}

func NewGameService(repo repository.GameRepository, logger *logrus.Logger) GameService {
	return &gameService{
		repo:   repo,
		logger: logger,
	}
}

func (s *gameService) CreateGame(ctx context.Context, game *models.Game) error {
	if game.Title == "" {
		return errors.New("le titre du jeu est obligatoire")
	}
	
	now := time.Now()
	game.CreatedAt = now
	game.UpdatedAt = now
	
	s.logger.WithFields(logrus.Fields{
		"title": game.Title,
	}).Info("Création d'un nouveau jeu")
	
	return s.repo.Create(ctx, game)
}

func (s *gameService) GetGameByID(ctx context.Context, id uint) (*models.Game, error) {
	s.logger.WithField("id", id).Info("Récupération d'un jeu")
	return s.repo.GetByID(ctx, id)
}

func (s *gameService) UpdateGame(ctx context.Context, game *models.Game) error {
	existingGame, err := s.repo.GetByID(ctx, game.ID)
	if err != nil {
		return err
	}
	
	game.UpdatedAt = time.Now()
	game.CreatedAt = existingGame.CreatedAt 
	
	s.logger.WithFields(logrus.Fields{
		"id":    game.ID,
		"title": game.Title,
	}).Info("Mise à jour d'un jeu")
	
	return s.repo.Update(ctx, game)
}

func (s *gameService) DeleteGame(ctx context.Context, id uint) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	
	s.logger.WithField("id", id).Info("Suppression d'un jeu")
	
	return s.repo.Delete(ctx, id)
}

func (s *gameService) ListGames(ctx context.Context, filter *models.GameFilter) (*models.GameResponse, error) {
	if filter == nil {
		filter = &models.GameFilter{
			Page:     1,
			PageSize: 10,
		}
	}
	
	s.logger.WithFields(logrus.Fields{
		"page":      filter.Page,
		"page_size": filter.PageSize,
		"title":     filter.Title,
	}).Info("Recherche de jeux")
	
	return s.repo.List(ctx, filter)
}

func (s *gameService) CreateGenre(ctx context.Context, genre *models.Genre) error {
	if genre.Name == "" {
		return errors.New("le nom du genre est obligatoire")
	}
	
	s.logger.WithField("name", genre.Name).Info("Création d'un nouveau genre")
	
	return s.repo.CreateGenre(ctx, genre)
}

func (s *gameService) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	s.logger.Info("Récupération de tous les genres")
	return s.repo.GetAllGenres(ctx)
}

func (s *gameService) CreatePlatform(ctx context.Context, platform *models.Platform) error {
	if platform.Name == "" {
		return errors.New("le nom de la plateforme est obligatoire")
	}
	
	s.logger.WithField("name", platform.Name).Info("Création d'une nouvelle plateforme")
	
	return s.repo.CreatePlatform(ctx, platform)
}

func (s *gameService) GetAllPlatforms(ctx context.Context) ([]models.Platform, error) {
	s.logger.Info("Récupération de toutes les plateformes")
	return s.repo.GetAllPlatforms(ctx)
}
