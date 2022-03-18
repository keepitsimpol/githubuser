package main

import (
	_ "github.com/keepitsimpol/githubuser/docs"

	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/core/service"
	"github.com/keepitsimpol/githubuser/internal/infrastructure"
	"github.com/keepitsimpol/githubuser/internal/middleware"
	"github.com/keepitsimpol/githubuser/internal/presentation"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const GinMode = gin.ReleaseMode

// @title          GithubUser APIs
// @version        1.0.0
// @description    A service that provides Github user details
// @contact.name   Pol Torres
// @contact.email  apolinario.torresjr@gmail.com
// @host           localhost:8080
// @BasePath       /api/v1
func main() {
	githubClientImpl := infrastructure.New()
	githubServiceImpl := service.New(githubClientImpl)
	githubUserController := presentation.New(githubServiceImpl)

	router := gin.Default()
	gin.SetMode(GinMode)

	router.Use(middleware.AppendRequestID)

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("users")
		{
			users.POST("/github", githubUserController.GetUserAccountDetails)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
