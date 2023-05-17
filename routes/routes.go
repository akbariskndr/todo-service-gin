package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	var router = gin.Default()

	v1 := router.Group("v1")
	AddPingRoutes(v1)
	AddTodosRoutes(v1)

	return router
}
