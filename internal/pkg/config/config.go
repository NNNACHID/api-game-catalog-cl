package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/NNNACHID/api-game-catalog-cl/pkg/database"
)

type Config struct {
	Server   ServerConfig
	Database database.PostgresConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type LoggerConfig struct {
	Level string
}


func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")          
	v.SetConfigType("yaml")           
	v.AddConfigPath(configPath)        
	v.AddConfigPath(".")           
	v.AddConfigPath("./config")       
	v.AddConfigPath("./configs")      
	v.AddConfigPath("$HOME/.appname") 

	setDefaults(v)

	v.AutomaticEnv()

	err := v.ReadInConfig(); 
	if err != nil {
		fmt.Printf("Attention: Fichier de configuration non trouvé ou illisible: %v\n", err)
	}

	config := Config{}
	err = v.Unmarshal(&config) 
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la désérialisation de la configuration: %w", err)
	}

	return &config, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.readTimeout", "5s")
	v.SetDefault("server.writeTimeout", "10s")
	v.SetDefault("server.idleTimeout", "120s")

	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", "5432")
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "postgres")
	v.SetDefault("database.dbname", "gamecatalog")
	v.SetDefault("database.sslmode", "disable")

	v.SetDefault("logger.level", "info")
}

func ConfigureLogger(config LoggerConfig) *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	return logger
}
