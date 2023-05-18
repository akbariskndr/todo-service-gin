package middleware

import (
	auth_module "akbariskndr/todo-service-gin/modules/auth"
	"akbariskndr/todo-service-gin/modules/auth/controller/dto"
	"akbariskndr/todo-service-gin/modules/auth/entity"
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CreateAuthMiddleware() *jwt.GinJWTMiddleware {
	auth := auth_module.InitModule()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("DB_USERNAME")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.UserEntity); ok {
				return jwt.MapClaims{
					"id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return auth.UserService.FindOne(uint(claims["id"].(float64)))
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var payload dto.LoginDto
			if err := c.ShouldBind(&payload); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if user := auth.UserService.Login(&payload); user != nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*entity.UserEntity); ok && v != nil {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}
