// main.go
package main

import (
	"context"
	"log"
	"os"
	"strings"

	"go-gin/cmd"
	"go-gin/cmd/commands"
	"go-gin/internal/api/handlers"
	"go-gin/internal/api/middleware"
	"go-gin/internal/api/routes"
	"go-gin/internal/domain/repository"
	"go-gin/internal/infrastructure/database"
	"go-gin/internal/infrastructure/kafka"
	"go-gin/internal/infrastructure/kafka/consumers/email"
	"go-gin/internal/logger"
	"go-gin/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize logger
	err = logger.Initialize(logger.Config{
		Environment: os.Getenv("APP_ENV"),
		SentryDSN:   os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Register commands
	cmd.RootCmd.AddCommand(commands.GetSendRemindersCmd())

	// If arguments are provided, run as CLI
	if len(os.Args) > 1 {
		cmd.Execute()
		return
	}

	// Otherwise, run as web server
	runServer()
}

func runServer() {
	// Initialize database and Redis
	if err := database.InitAllDatabases(); err != nil {
		log.Fatal("Failed to initialize databases:", err)
	}

	// Get CRM database
	crmDB, err := database.GetDB("gin")
	if err != nil {
		log.Fatal(err)
	}

	// Get IMS database
	/*
		imsDB, err := database.GetDB("ims")
		if err != nil {
			log.Fatal(err)
		}*/

	// Initialize Kafka
	kafkaBrokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	producer := kafka.NewProducer(kafkaBrokers)
	defer producer.Close()

	// Start email consumer
	emailConsumer := email.NewEmailConsumer(kafkaBrokers)
	go emailConsumer.Start(context.Background())
	defer emailConsumer.Close()

	// Setup repositories, services, and handlers
	userRepo := repository.NewUserRepository(crmDB)
	taskRepo := repository.NewTaskRepository(crmDB)

	authService := service.NewAuthService(userRepo, database.RDB, producer)
	taskService := service.NewTaskService(taskRepo)

	authHandler := handlers.NewAuthHandler(authService)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()

	// Public routes
	routes.PublicRoutes(r, authHandler)

	// Protected routes
	authMiddleware := middleware.NewAuthMiddleware(authService)
	protected := r.Group("/api")
	protected.Use(authMiddleware.JWTAuth())
	routes.ProtectedRoutes(protected, taskHandler)

	logger.Info("Server starting", map[string]interface{}{
		"port": os.Getenv("PORT"),
		"env":  os.Getenv("APP_ENV"),
	})

	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port) // Required to allow access from host machine

}
