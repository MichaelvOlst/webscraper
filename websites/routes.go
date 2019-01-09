package websites

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes for website urls
func RegisterRoutes(r *gin.RouterGroup) {

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"data": "websites",
		})
	})
}
