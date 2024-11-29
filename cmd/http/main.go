package main

import (
	"github.com/xvrzhao/groove-scaffold/handler"
	"github.com/xvrzhao/groove-scaffold/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	r.Use(cors.New(corsConfig))

	api := r.Group("/api")
	{
		// Auth
		auth := api.Group("/auth")
		{
			authCtrl := handler.NewAuthCtrl()
			auth.POST("/login", authCtrl.Login)
			auth.PUT("/password", middleware.Auth, authCtrl.ChangePassword)
		}

		// Users
		user := api.Group("/users", middleware.Auth)
		{
			userCtrl := handler.NewUserCtrl()
			user.GET("", userCtrl.Page)
			user.POST("", userCtrl.Create)
			user.PUT("/:id", userCtrl.Update)
			user.DELETE("/:id", userCtrl.Delete)
		}
	}

	r.Run(":80")
}
