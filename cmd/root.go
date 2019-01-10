package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"michaelvanolst.nl/scraper/datastore"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initApp)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.config.json)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName(".config")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			panic(err)
		}
	} else {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// App ...
type App struct {
	database datastore.Datastore
}

var app *App

func initApp() {

	db, err := datastore.New()
	if err != nil {
		fmt.Errorf("Error occured: connecting DB: %v", err)
		return
	}
	app = &App{
		database: db,
	}
}

var rootCmd = &cobra.Command{
	Use:   "Scraper",
	Short: "It scrapes websites.",
}

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
