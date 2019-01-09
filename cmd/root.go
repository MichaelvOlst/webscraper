package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"michaelvanolst.nl/scraper/storage"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	// cobra.OnInitialize(initDatabase)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.config.json)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
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

var store *storage.Storage

func initDatabase() {
	storage, err := storage.New()
	if err != nil {
		log.Printf("Error starting db %v", err)
		return
	}
	store = storage
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
