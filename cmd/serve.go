package cmd

import (
	"fmt"
	"net/http"
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

		server := &http.Server{
			Addr:         addr,
			Handler:      h,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		logrus.Infof("Server is now listening on http://%s", server.Addr)
		logrus.Fatal(server.ListenAndServe())
	},
}
