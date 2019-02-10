package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"michaelvanolst.nl/scraper/api"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the webserver",
	Run: func(cmd *cobra.Command, args []string) {
		var h http.Handler
		a := api.New(app.database)
		h = a.Routes()

		addr := fmt.Sprintf("%s:%s", app.config.Server.Host, app.config.Server.Port)

		app.server = &http.Server{
			Addr:         addr,
			Handler:      h,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		sigs := make(chan os.Signal, 1)
		done := make(chan struct{}, 1)
		signal.Notify(sigs, os.Interrupt, os.Kill)

		go func() {
			sig := <-sigs
			fmt.Println()
			fmt.Println(sig)
			done <- struct{}{}

			if err := app.server.Shutdown(context.Background()); err != nil {
				log.Printf("Unable to shut down server: %v", err)
			} else {
				log.Println("Server stopped")
			}

		}()

		log.Printf("Starting HTTP Server. Listening at http://%s", app.server.Addr)
		if err := app.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("%v", err)
		} else {
			log.Println("Server closed!")
		}

		app.cronjob.Close()

		err := app.database.Close()
		if err != nil {
			logrus.Errorf("Error closing database: %v", err)
			os.Exit(1)
		}

		// fmt.Println("awaiting signal")
		// <-done
		// fmt.Println("exiting")
	},
}
