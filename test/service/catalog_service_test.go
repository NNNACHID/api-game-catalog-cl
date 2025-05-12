package service

import (
	"context"
	"errors"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/models"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/service"
	//"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/repository"
)

type MockGameRepository struct {
	mock.Mock
}

func (m *MockGameRepository) Create(ctx context.Context, game *models.Game) error {
	args := m.Called(ctx, game)
	return args.Error(0)
}

func (m *MockGameRepository) GetByID(ctx context.Context, id uint) (*models.Game, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Game), args.Error(1)
}

func (m *MockGameRepository) Update(ctx context.Context, game *models.Game) error {
	args := m.Called(ctx, game)
	return args.Error(0)
}

func (m *MockGameRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockGameRepository) List(ctx context.Context, filter *models.GameFilter) (*models.GameResponse, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GameResponse), args.Error(1)
}

func (m *MockGameRepository) CreateGenre(ctx context.Context, genre *models.Genre) error {
	args := m.Called(ctx, genre)
	return args.Error(0)
}

func (m *MockGameRepository) GetAllGenres(ctx context.Context) ([]models.Genre, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Genre), args.Error(1)
}

func (m *MockGameRepository) CreatePlatform(ctx context.Context, platform *models.Platform) error {
	args := m.Called(ctx, platform)
	return args.Error(0)
}

func (m *MockGameRepository) GetAllPlatforms(ctx context.Context) ([]models.Platform, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Platform), args.Error(1)
}

func setupTest() (*MockGameRepository, service.GameService) {
	mockRepo := new(MockGameRepository)
	logger := logrus.New()
	logger.SetOutput(logrus.StandardLogger().Out)
	service := service.NewGameService(mockRepo, logger)

	return mockRepo, service
}

func TestCreateGame(t *testing.T) {
	t.Run("succès création jeu", func(t *testing.T) {
		mockRepo, service := setupTest()
		ctx := context.Background()
		game := &models.Game{
			Title:       "Test Game",
			Description: "Test Description",
		}
		mockRepo.On("Create", ctx, mock.MatchedBy(func(g *models.Game) bool {
			return g.Title == game.Title && g.Description == game.Description
		})).Return(nil)

		err := service.CreateGame(ctx, game)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("échec création jeu - titre manquant", func(t *testing.T) {
		mockRepo, service := setupTest()
		ctx := context.Background()
		game := &models.Game{
			Description: "Test Description",
		}

		err := service.CreateGame(ctx, game)

		assert.Error(t, err)
		assert.Equal(t, "le titre du jeu est obligatoire", err.Error())
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("échec création jeu - erreur repository", func(t *testing.T) {
		mockRepo, service := setupTest()
		ctx := context.Background()
		game := &models.Game{
			Title:       "Test Game",
			Description: "Test Description",
		}
		expectedErr := errors.New("erreur de base de données")
		mockRepo.On("Create", ctx, mock.MatchedBy(func(g *models.Game) bool {
			return g.Title == game.Title && g.Description == game.Description
		})).Return(expectedErr)

		err := service.CreateGame(ctx, game)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

// func TestGetGameByID(t *testing.T) {
// 	t.Run("succès récupération jeu", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		expectedGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "Test Game",
// 			Description: "Test Description",
// 			CreatedAt:   time.Now(),
// 			UpdatedAt:   time.Now(),
// 		}
// 		mockRepo.On("GetByID", ctx, gameID).Return(expectedGame, nil)

// 		// Act
// 		game, err := service.GetGameByID(ctx, gameID)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedGame, game)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec récupération jeu - jeu non trouvé", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		expectedErr := errors.New("jeu non trouvé")
// 		mockRepo.On("GetByID", ctx, gameID).Return(nil, expectedErr)

// 		// Act
// 		game, err := service.GetGameByID(ctx, gameID)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Nil(t, game)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestUpdateGame(t *testing.T) {
// 	t.Run("succès mise à jour jeu", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		createdTime := time.Now().Add(-24 * time.Hour)
// 		existingGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "Old Title",
// 			Description: "Old Description",
// 			CreatedAt:   createdTime,
// 			UpdatedAt:   createdTime,
// 		}
// 		updatedGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "New Title",
// 			Description: "New Description",
// 		}
// 		mockRepo.On("GetByID", ctx, gameID).Return(existingGame, nil)
// 		mockRepo.On("Update", ctx, mock.MatchedBy(func(g *models.Game) bool {
// 			return g.ID == gameID && g.Title == updatedGame.Title && g.CreatedAt == createdTime && !g.UpdatedAt.Equal(createdTime)
// 		})).Return(nil)

// 		// Act
// 		err := service.UpdateGame(ctx, updatedGame)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, createdTime, updatedGame.CreatedAt)
// 		assert.NotEqual(t, createdTime, updatedGame.UpdatedAt)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec mise à jour jeu - jeu non trouvé", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		updatedGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "New Title",
// 			Description: "New Description",
// 		}
// 		expectedErr := errors.New("jeu non trouvé")
// 		mockRepo.On("GetByID", ctx, gameID).Return(nil, expectedErr)

// 		// Act
// 		err := service.UpdateGame(ctx, updatedGame)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertNotCalled(t, "Update")
// 	})

// 	t.Run("échec mise à jour jeu - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		createdTime := time.Now().Add(-24 * time.Hour)
// 		existingGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "Old Title",
// 			Description: "Old Description",
// 			CreatedAt:   createdTime,
// 			UpdatedAt:   createdTime,
// 		}
// 		updatedGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "New Title",
// 			Description: "New Description",
// 		}
// 		expectedErr := errors.New("erreur de mise à jour")
// 		mockRepo.On("GetByID", ctx, gameID).Return(existingGame, nil)
// 		mockRepo.On("Update", ctx, mock.MatchedBy(func(g *models.Game) bool {
// 			return g.ID == gameID && g.Title == updatedGame.Title && g.CreatedAt == createdTime && !g.UpdatedAt.Equal(createdTime)
// 		})).Return(expectedErr)

// 		// Act
// 		err := service.UpdateGame(ctx, updatedGame)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestDeleteGame(t *testing.T) {
// 	t.Run("succès suppression jeu", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		existingGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "Test Game",
// 			Description: "Test Description",
// 		}
// 		mockRepo.On("GetByID", ctx, gameID).Return(existingGame, nil)
// 		mockRepo.On("Delete", ctx, gameID).Return(nil)

// 		// Act
// 		err := service.DeleteGame(ctx, gameID)

// 		// Assert
// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec suppression jeu - jeu non trouvé", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		expectedErr := errors.New("jeu non trouvé")
// 		mockRepo.On("GetByID", ctx, gameID).Return(nil, expectedErr)

// 		// Act
// 		err := service.DeleteGame(ctx, gameID)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertNotCalled(t, "Delete")
// 	})

// 	t.Run("échec suppression jeu - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		gameID := uint(1)
// 		existingGame := &models.Game{
// 			ID:          gameID,
// 			Title:       "Test Game",
// 			Description: "Test Description",
// 		}
// 		expectedErr := errors.New("erreur de suppression")
// 		mockRepo.On("GetByID", ctx, gameID).Return(existingGame, nil)
// 		mockRepo.On("Delete", ctx, gameID).Return(expectedErr)

// 		// Act
// 		err := service.DeleteGame(ctx, gameID)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestListGames(t *testing.T) {
// 	t.Run("succès liste jeux avec filtre", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		filter := &models.GameFilter{
// 			Page:     2,
// 			PageSize: 15,
// 			Title:    "Test",
// 		}
// 		expectedResponse := &models.GameResponse{
// 			Games: []models.Game{
// 				{ID: 1, Title: "Test Game 1"},
// 				{ID: 2, Title: "Test Game 2"},
// 			},
// 			TotalItems:  2,
// 			TotalPages:  1,
// 			CurrentPage: 2,
// 		}
// 		mockRepo.On("List", ctx, filter).Return(expectedResponse, nil)

// 		// Act
// 		response, err := service.ListGames(ctx, filter)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResponse, response)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("succès liste jeux sans filtre", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		expectedResponse := &models.GameResponse{
// 			Games: []models.Game{
// 				{ID: 1, Title: "Game 1"},
// 				{ID: 2, Title: "Game 2"},
// 			},
// 			TotalItems:  2,
// 			TotalPages:  1,
// 			CurrentPage: 1,
// 		}
// 		mockRepo.On("List", ctx, mock.MatchedBy(func(f *models.GameFilter) bool {
// 			return f.Page == 1 && f.PageSize == 10 && f.Title == ""
// 		})).Return(expectedResponse, nil)

// 		// Act
// 		response, err := service.ListGames(ctx, nil)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedResponse, response)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec liste jeux - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		filter := &models.GameFilter{
// 			Page:     1,
// 			PageSize: 10,
// 		}
// 		expectedErr := errors.New("erreur de liste")
// 		mockRepo.On("List", ctx, filter).Return(nil, expectedErr)

// 		// Act
// 		response, err := service.ListGames(ctx, filter)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Nil(t, response)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestCreateGenre(t *testing.T) {
// 	t.Run("succès création genre", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		genre := &models.Genre{
// 			Name: "Action",
// 		}
// 		mockRepo.On("CreateGenre", ctx, genre).Return(nil)

// 		// Act
// 		err := service.CreateGenre(ctx, genre)

// 		// Assert
// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec création genre - nom manquant", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		genre := &models.Genre{}

// 		// Act
// 		err := service.CreateGenre(ctx, genre)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, "le nom du genre est obligatoire", err.Error())
// 		mockRepo.AssertNotCalled(t, "CreateGenre")
// 	})

// 	t.Run("échec création genre - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		genre := &models.Genre{
// 			Name: "Action",
// 		}
// 		expectedErr := errors.New("erreur de base de données")
// 		mockRepo.On("CreateGenre", ctx, genre).Return(expectedErr)

// 		// Act
// 		err := service.CreateGenre(ctx, genre)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestGetAllGenres(t *testing.T) {
// 	t.Run("succès récupération genres", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		expectedGenres := []models.Genre{
// 			{ID: 1, Name: "Action"},
// 			{ID: 2, Name: "Adventure"},
// 		}
// 		mockRepo.On("GetAllGenres", ctx).Return(expectedGenres, nil)

// 		// Act
// 		genres, err := service.GetAllGenres(ctx)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedGenres, genres)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec récupération genres - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		expectedErr := errors.New("erreur de base de données")
// 		mockRepo.On("GetAllGenres", ctx).Return([]models.Genre{}, expectedErr)

// 		// Act
// 		genres, err := service.GetAllGenres(ctx)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		assert.Empty(t, genres)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestCreatePlatform(t *testing.T) {
// 	t.Run("succès création plateforme", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		platform := &models.Platform{
// 			Name: "PlayStation 5",
// 		}
// 		mockRepo.On("CreatePlatform", ctx, platform).Return(nil)

// 		// Act
// 		err := service.CreatePlatform(ctx, platform)

// 		// Assert
// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec création plateforme - nom manquant", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		platform := &models.Platform{}

// 		// Act
// 		err := service.CreatePlatform(ctx, platform)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, "le nom de la plateforme est obligatoire", err.Error())
// 		mockRepo.AssertNotCalled(t, "CreatePlatform")
// 	})

// 	t.Run("échec création plateforme - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		platform := &models.Platform{
// 			Name: "PlayStation 5",
// 		}
// 		expectedErr := errors.New("erreur de base de données")
// 		mockRepo.On("CreatePlatform", ctx, platform).Return(expectedErr)

// 		// Act
// 		err := service.CreatePlatform(ctx, platform)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		mockRepo.AssertExpectations(t)
// 	})
// }

// func TestGetAllPlatforms(t *testing.T) {
// 	t.Run("succès récupération plateformes", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		expectedPlatforms := []models.Platform{
// 			{ID: 1, Name: "PlayStation 5"},
// 			{ID: 2, Name: "Xbox Series X"},
// 			{ID: 3, Name: "Nintendo Switch"},
// 		}
// 		mockRepo.On("GetAllPlatforms", ctx).Return(expectedPlatforms, nil)

// 		// Act
// 		platforms, err := service.GetAllPlatforms(ctx)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, expectedPlatforms, platforms)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("échec récupération plateformes - erreur repository", func(t *testing.T) {
// 		// Arrange
// 		mockRepo, service := setupTest()
// 		ctx := context.Background()
// 		expectedErr := errors.New("erreur de base de données")
// 		mockRepo.On("GetAllPlatforms", ctx).Return([]models.Platform{}, expectedErr)

// 		// Act
// 		platforms, err := service.GetAllPlatforms(ctx)

// 		// Assert
// 		assert.Error(t, err)
// 		assert.Equal(t, expectedErr, err)
// 		assert.Empty(t, platforms)
// 		mockRepo.AssertExpectations(t)
// 	})
// }