package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/smartnotes/user-service/internal/handlers"
	"github.com/smartnotes/user-service/internal/middleware"
	"github.com/smartnotes/user-service/internal/repositories"
)

func main() {
	// Load environment variables
	_ = godotenv.Load() // Просто попытка загрузить .env, но не критично если его нет

	// Initialize MongoDB connection
	repo, err := repositories.NewMongoDBRepository()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(repo)
	bookHandler := handlers.NewBookHandler(repo)

	// Create router
	router := gin.Default()

	// Log all incoming requests (including OPTIONS)
	router.Use(func(c *gin.Context) {
		log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	router.POST("/api/auth/register", authHandler.Register)
	router.POST("/api/auth/login", authHandler.Login)

	// Protected routes
	auth := router.Group("/api/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/me", authHandler.GetCurrentUser)
	}

	// Book routes
	books := router.Group("/api/books")
	books.Use(middleware.AuthMiddleware())
	{
		books.POST("", bookHandler.CreateBook)
		books.GET("", bookHandler.ListBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("User Service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
