package main

import (
	"log"
	"rogeriods/contact-api/internal/database"
	"rogeriods/contact-api/internal/handler"
	"rogeriods/contact-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load .env variables
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	db := database.Init()

	r := gin.Default()

	authHandler := handler.NewAuthHandler(db)
	contactHandler := handler.NewContactHandler(db)

	// Public
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Protected
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/contacts", contactHandler.GetContacts)
		auth.GET("/contacts/:id", contactHandler.GetByID)
		auth.POST("/contacts", contactHandler.Create)
		auth.PUT("/contacts/:id", contactHandler.Update)
		auth.DELETE("/contacts/:id", contactHandler.Delete)
	}

	r.Run(":8080")
}