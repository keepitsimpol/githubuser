package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keepitsimpol/githubuser/internal/presentation"
)

func main() {
	githubUserController := presentation.New()
	router := gin.Default()
	router.GET("/user", githubUserController.GetGethubUser())
	router.Run()
}
