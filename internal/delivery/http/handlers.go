package http

import (
	"net/http"
	"strconv"

	"github.com/NNNACHID/api-game-catalog-cl/internal/models"
	"github.com/NNNACHID/api-game-catalog-cl/internal/service"
	"github.com/gin-gonic/gin"
	//"github.com/golangci/golangci-lint/pkg/golinters/iface"
	"github.com/sirupsen/logrus"
)

type GameHandler struct {
	service service.GameService
	logger  *logrus.Logger
}

func NewGameHandler(service service.GameService, logger *logrus.Logger) *GameHandler {
	return &GameHandler{
		service: service,
		logger:  logger,
	}
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	var game models.Game
	
	err := c.ShouldBindJSON(&game)
	if err != nil {
		h.logger.WithError(err).Error("Error deserializing request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})

		return
	}
	
	err = h.service.CreateGame(c.Request.Context(), &game)
	if err != nil {
		h.logger.WithError(err).Error("Error creating game")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusCreated, game)
}

func (h *GameHandler) GetGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid game ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})

		return
	}
	
	game, err := h.service.GetGameByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Error retrieving game")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) UpdateGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid game ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})

		return
	}
	
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		h.logger.WithError(err).Error("Error deserializing request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})

		return
	}
	
	game.ID = uint(id)
	
	if err := h.service.UpdateGame(c.Request.Context(), &game); err != nil {
		h.logger.WithError(err).Error("Error updating the game")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) DeleteGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("Invalid game ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid game ID"})

		return
	}
	
	err = h.service.DeleteGame(c.Request.Context(), uint(id)) 
	if err != nil {
		h.logger.WithError(err).Error("Error deleting game")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}

func (h *GameHandler) ListGames(c *gin.Context) {
	var filter models.GameFilter
	
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		h.logger.WithError(err).Error("Error deserializing request params")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query params"})

		return
	}
	
	games, err := h.service.ListGames(c.Request.Context(), &filter)
	if err != nil {
		h.logger.WithError(err).Error("Error retrieving games")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) CreateGenre(c *gin.Context) {
	var genre models.Genre
	
	if err := c.ShouldBindJSON(&genre); err != nil {
		h.logger.WithError(err).Error("Error deserializing request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query format"})

		return
	}
	
	if err := h.service.CreateGenre(c.Request.Context(), &genre); err != nil {
		h.logger.WithError(err).Error("Error creating platform")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusCreated, genre)
}

func (h *GameHandler) GetAllGenres(c *gin.Context) {
	genres, err := h.service.GetAllGenres(c.Request.Context())
	if err != nil {
		h.logger.WithError(err).Error("Error retrieving genres")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, genres)
}

func (h *GameHandler) CreatePlatform(c *gin.Context) {
	platform := models.Platform{}
	
	err := c.ShouldBindJSON(&platform)
	if err != nil {
		h.logger.WithError(err).Error("Error deserializing request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query format"})

		return
	}
	
	err = h.service.CreatePlatform(c.Request.Context(), &platform)
	if err != nil {
		h.logger.WithError(err).Error("Error creating platform")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusCreated, platform)
}

func (h *GameHandler) GetAllPlatforms(c *gin.Context) {
	platforms, err := h.service.GetAllPlatforms(c.Request.Context())
	if err != nil {
		h.logger.WithError(err).Error("Error retrieving platforms")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, platforms)
}
