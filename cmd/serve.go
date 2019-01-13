package cmd

import (
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
		a := api.New(app.database)
		a.Start()

		// fmt.Println(app.config.Server.Port)

	},
}
