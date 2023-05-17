package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddPingRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/ping")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{"message": "pong"})
	})
}
