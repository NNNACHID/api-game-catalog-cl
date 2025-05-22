package migrations

import (
	"github.com/sirupsen/logrus"
	"github.com/NNNACHID/api-game-catalog-cl/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB, logger *logrus.Logger) error {
	logger.Info("Début des migrations de base de données")

	models := []interface{}{
		&models.Game{},
		&models.Genre{},
		&models.Platform{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			logger.WithError(err).Errorf("Erreur lors de la migration pour %T", model)

			return err
		}

		logger.Infof("Migration réussie pour %T", model)
	}

	logger.Info("Migrations de base de données terminées avec succès")
	return nil
}

func SeedData(db *gorm.DB, logger *logrus.Logger) error {
	var count int64
	db.Model(&models.Genre{}).Count(&count)
	
	if count > 0 {
		logger.Info("Des données existent déjà, pas de seeding nécessaire")
		return nil
	}
	
	logger.Info("Initialisation des données de test")
	
	genres := []models.Genre{
		{Name: "Action"},
		{Name: "Aventure"},
		{Name: "RPG"},
		{Name: "FPS"},
		{Name: "Stratégie"},
		{Name: "Simulation"},
		{Name: "Sport"},
		{Name: "Course"},
		{Name: "Puzzle"},
		{Name: "Plateforme"},
	}
	
	err := db.CreateInBatches(genres, len(genres)).Error 
	if err != nil {
		logger.WithError(err).Error("Erreur lors de la création des genres")
		return err
	}
	
	platforms := []models.Platform{
		{Name: "PC"},
		{Name: "PlayStation 5"},
		{Name: "PlayStation 4"},
		{Name: "Xbox Series X"},
		{Name: "Xbox One"},
		{Name: "Nintendo Switch"},
		{Name: "Mobile"},
	}
	
	err = db.CreateInBatches(platforms, len(platforms)).Error 
	if err != nil {
		logger.WithError(err).Error("Erreur lors de la création des plateformes")
		return err
	}

	logger.Info("Données de test initialisées avec succès")
	return nil
}