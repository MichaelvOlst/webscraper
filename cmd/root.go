package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"michaelvanolst.nl/scraper/config"
	"michaelvanolst.nl/scraper/datastore"
)

var configFile string

func init() {
	cobra.OnInitialize(initApp)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is .env)")
}

// App ...
type App struct {
	database datastore.Datastore
	config   *config.Config
}

var app *App

func initApp() {

	err := config.Load(configFile)
	if err != nil {
		log.Errorf("Cannot load config: %v", err)
		os.Exit(1)
	}

	config, err := config.Parse()
	if err != nil {
		log.Errorf("Cannot load config: %v", err)
		os.Exit(1)
	}

	db, err := datastore.New()
	if err != nil {
		log.Errorf("Cannot init db error: %v", err)
		os.Exit(1)
	}

	app = &App{
		database: db,
		config:   config,
	}

}

var rootCmd = &cobra.Command{
	Use:   "Scraper",
	Short: "It scrapes websites.",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		err := app.database.Close()
		if err != nil {
			log.Errorf("Error closing database: %v", err)
			os.Exit(1)
		}
	},
}

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("Error executing command: %v", err)
		os.Exit(1)
	}
}
