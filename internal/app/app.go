package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"template/internal/app/config"
	"template/internal/app/connections"
	"template/internal/app/store"
	"template/internal/services/course"
)

func Run(filenames ...string) {
	// Load configuration
	cfg, err := config.New(filenames...)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize connections
	conns, err := connections.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create connections: %v", err)
	}

	// Initialize store and services
	st := store.NewRepositoryStore(conns)
	clients := store.NewClientStore(conns)
	courseService := course.New(st, clients)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/v1/courses", courseService.CreateCourse)
	e.GET("/v1/courses", courseService.ListCourses)
	e.GET("/v1/courses/:id", courseService.GetCourse)
	e.DELETE("/v1/courses/:id", courseService.DeleteCourse)
	e.POST("/v1/courses/search", courseService.SearchCourses)

	// Start server
	go func() {
		if err := e.Start(cfg.HTTP.Addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
