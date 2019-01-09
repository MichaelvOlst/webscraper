package server

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"michaelvanolst.nl/scraper/websites"
)

// Server holds the router, db and maybe more
type Server struct {
	router *gin.Engine
}

// New inits the server
func New() *Server {

	return &Server{
		router: gin.Default(),
	}
}

// Start the server
func (s *Server) Start() {
	s.router.Use(static.Serve("/", static.LocalFile("public", true)))
	api := s.router.Group("/api")
	websites.RegisterRoutes(api.Group("/websites"))

	address := fmt.Sprintf("%s:%s", viper.Get("server.address"), viper.Get("server.port"))
	s.router.Run(address)
}
