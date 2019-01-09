package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"michaelvanolst.nl/scraper/server"
)

var port string
var address string

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "The port where the webserver needs to run on")
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))

	serveCmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1", "The address where the webserver needs to run on")
	viper.BindPFlag("server.address", serveCmd.Flags().Lookup("address"))
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the webserver",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New()
		s.Start()
	},
}
