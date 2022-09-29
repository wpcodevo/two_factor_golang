package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/two_factor_golang/controllers"
	"github.com/wpcodevo/two_factor_golang/models"
	"github.com/wpcodevo/two_factor_golang/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	AuthController = controllers.NewAuthController(DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Two-Factor Authentication with Golang"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	log.Fatal(server.Run(":8000"))
}
