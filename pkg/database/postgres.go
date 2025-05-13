package database

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(config PostgresConfig, log *logrus.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	gormLogger := logger.New(
		&logrusWriter{log},
		logger.Config{
			SlowThreshold:             time.Second, 
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,   
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("erreur de connexion à PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'accès au pool de connexions SQL: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info("Connexion à la base de données PostgreSQL établie avec succès")
	return db, nil
}

type logrusWriter struct {
	logger *logrus.Logger
}

func (l *logrusWriter) Printf(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}
