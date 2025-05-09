package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/repository"
	"github.com/NNNACHID/api-game-catalog-cl/internal/catalog/service"
	catalogHTTP "github.com/NNNACHID/api-game-catalog-cl/internal/catalog/delivery/http"
	"github.com/NNNACHID/api-game-catalog-cl/internal/pkg/config"
	"github.com/NNNACHID/api-game-catalog-cl/internal/pkg/migrations"
	"github.com/NNNACHID/api-game-catalog-cl/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("Erreur lors du chargement de la configuration: %v\n", err)
		os.Exit(1)
	}

	logger := config.ConfigureLogger(cfg.Logger)
	logger.Info("Démarrage du service de catalogue")

	db, err := database.NewPostgresConnection(cfg.Database, logger)
	if err != nil {
		logger.WithError(err).Fatal("Impossible de se connecter à la base de données")
	}

	err = migrations.RunMigrations(db, logger) 
	if err != nil {
		logger.WithError(err).Fatal("Erreur lors des migrations")
	}

	err = migrations.SeedData(db, logger) 
	if err != nil {
		logger.WithError(err).Fatal("Erreur lors de l'initialisation des données de test")
	}

	gameRepo := repository.NewPostgresGameRepository(db)
	gameService := service.NewGameService(gameRepo, logger)
	gameHandler := catalogHTTP.NewGameHandler(gameService, logger)

	router := gin.New()
	router.Use(gin.Recovery())
	
	router.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		
		c.Next()
		
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		
		logger.WithFields(logrus.Fields{
			"status":     statusCode,
			"latency":    latency,
			"path":       path,
			"method":     c.Request.Method,
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Info("Requête HTTP")
	})

	gameHandler.RegisterRoutes(router)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	go func() {
		logger.Infof("Serveur HTTP démarré sur le port %s", cfg.Server.Port)
		err := server.ListenAndServe()
		if err != nil{
			logger.WithError(err).Fatal("Erreur lors du démarrage du serveur HTTP")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Arrêt du serveur en cours...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		logger.WithError(err).Fatal("Erreur lors de l'arrêt du serveur")
	}

	logger.Info("Serveur arrêté avec succès")
}
