package http

import (
	"github.com/gin-gonic/gin"
)

func (h *GameHandler) RegisterRoutes(router *gin.Engine) {
	catalog := router.Group("/api/v1/catalog")
	{
		catalog.POST("/games", h.CreateGame)
		catalog.GET("/games/:id", h.GetGame)
		catalog.PUT("/games/:id", h.UpdateGame)
		catalog.DELETE("/games/:id", h.DeleteGame)
		catalog.GET("/games", h.ListGames)
		
		catalog.POST("/genres", h.CreateGenre)
		catalog.GET("/genres", h.GetAllGenres)
		
		catalog.POST("/platforms", h.CreatePlatform)
		catalog.GET("/platforms", h.GetAllPlatforms)
	}
}