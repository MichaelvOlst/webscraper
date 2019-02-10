package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"michaelvanolst.nl/scraper/cronjob"
	"michaelvanolst.nl/scraper/websites"

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

	db, err := datastore.New(config.Database)
	if err != nil {
		log.Errorf("Cannot init db error: %v", err)
		os.Exit(1)
	}

	app = &App{
		database: db,
		config:   config,
	}

	c := cronjob.New(app.database)
	c.AddJob(60*60, func() {
		websites.Scrape(app.database)
	})
	c.Start()

}

var rootCmd = &cobra.Command{
	Use:   "Scraper",
	Short: "It scrapes websites.",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// err := app.database.Close()
		// if err != nil {
		// 	log.Errorf("Error closing database: %v", err)
		// 	os.Exit(1)
		// }
	},
}

// Execute ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("Error executing command: %v", err)
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- struct{}{}
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
	err := app.database.Close()
	if err != nil {
		log.Errorf("Error closing database: %v", err)
		os.Exit(1)
	}
}

// func runCronjob() {

// 	ticker := time.NewTicker(5 * time.Second)

// 	sigs := make(chan os.Signal, 1)
// 	done := make(chan struct{}, 1)

// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

// 	go func() {
// 		sig := <-sigs
// 		fmt.Println()
// 		fmt.Println(sig)
// 		done <- struct{}{}
// 	}()

// 	go func() {
// 		// var count int = 0
// 		for {
// 			select {
// 			case <-ticker.C:
// 				websites.Scrape(app.database)
// 			case <-done:
// 				fmt.Printf("Done counting.. \n")
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// 	fmt.Println("awaiting signal")
// 	<-done
// 	fmt.Println("exiting")
// }
