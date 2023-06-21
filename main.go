package main

import (
	"go-project/config"
	"go-project/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	config.Connect()

	router := gin.Default()

	// Routes
	router.POST("/users", controllers.CreateUser)

	router.GET("/users", controllers.GetUsers)

	router.GET("/users/:id", controllers.GetUser)

	router.PUT("/users/:id", controllers.UpdateUser)

	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run(":" + os.Getenv("PORT"))
}
