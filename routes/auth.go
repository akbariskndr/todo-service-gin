package routes

import (
	"akbariskndr/todo-service-gin/modules/auth/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(rg *gin.RouterGroup, controller *controller.AuthController, authMiddleware *jwt.GinJWTMiddleware) {
	router := rg.Group("/auth")

	router.GET("/me", authMiddleware.MiddlewareFunc(), controller.Me)
	router.POST("/register", controller.Register)
	router.POST("/login", authMiddleware.LoginHandler)
	router.GET("/refresh-token", authMiddleware.RefreshHandler)
	router.PATCH("/change-password", authMiddleware.MiddlewareFunc(), controller.ChangePassword)
}
