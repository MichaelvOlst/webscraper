package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"michaelvanolst.nl/scraper/api"
	"michaelvanolst.nl/scraper/datastore"
	"michaelvanolst.nl/scraper/email"
)

// Config hold the config for this app
type Config struct {
	Database *datastore.Config
	Server   *api.Config
	Email    *email.Config
}

// Load loads the config file
func Load(file string) error {

	if file == "" {
		log.Warn("Using default config")
		return nil
	}

	absFile, _ := filepath.Abs(file)
	_, err := os.Stat(absFile)
	fileNotExists := os.IsNotExist(err)

	if fileNotExists {
		return errors.New("Error reading configuration. File " + file + " does not exist.")
	}

	log.Printf("Configuration file: %s", absFile)

	// read file into env values
	err = godotenv.Load(absFile)
	if err != nil {
		return err
	}

	return nil
}

// Parse handles the config
func Parse() (*Config, error) {
	var cfg Config

	// with config file loaded into env values, we can now parse env into our config struct
	err := envconfig.Process("app", &cfg)
	if err != nil {
		return nil, err
	}

	// alias sqlite to sqlite3
	if cfg.Database.Driver == "sqlite" {
		cfg.Database.Driver = "sqlite3"
	}

	// use absolute path to sqlite3 database
	if cfg.Database.Driver == "sqlite3" {
		cfg.Database.Name, _ = filepath.Abs(cfg.Database.Name)
	}

	return &cfg, nil
}
