package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/models"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/service"
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
	
	if err := c.ShouldBindJSON(&game); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la désérialisation de la requête")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de requête invalide"})

		return
	}
	
	if err := h.service.CreateGame(c.Request.Context(), &game); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la création du jeu")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusCreated, game)
}

func (h *GameHandler) GetGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		h.logger.WithError(err).Error("ID de jeu invalide")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de jeu invalide"})

		return
	}
	
	game, err := h.service.GetGameByID(c.Request.Context(), uint(id))
	if err != nil {
		h.logger.WithError(err).Error("Erreur lors de la récupération du jeu")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) UpdateGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("ID de jeu invalide")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de jeu invalide"})

		return
	}
	
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la désérialisation de la requête")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de requête invalide"})

		return
	}
	
	game.ID = uint(id)
	
	if err := h.service.UpdateGame(c.Request.Context(), &game); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la mise à jour du jeu")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) DeleteGame(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.WithError(err).Error("ID de jeu invalide")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de jeu invalide"})

		return
	}
	
	if err := h.service.DeleteGame(c.Request.Context(), uint(id)); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la suppression du jeu")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Jeu supprimé avec succès"})
}

func (h *GameHandler) ListGames(c *gin.Context) {
	var filter models.GameFilter
	
	if err := c.ShouldBindQuery(&filter); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la désérialisation des paramètres de requête")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Paramètres de requête invalides"})

		return
	}
	
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 {
		filter.PageSize = 10
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
		h.logger.WithError(err).Error("Erreur lors de la désérialisation de la requête")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de requête invalide"})

		return
	}
	
	if err := h.service.CreateGenre(c.Request.Context(), &genre); err != nil {
		h.logger.WithError(err).Error("Erreur lors de la création du genre")
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
		h.logger.WithError(err).Error("Erreur lors de la désérialisation de la requête")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format de requête invalide"})

		return
	}
	
	err = h.service.CreatePlatform(c.Request.Context(), &platform)
	if err != nil {
		h.logger.WithError(err).Error("Erreur lors de la création du genre")
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