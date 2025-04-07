package main

import (
	"log"

	"go-slotcars/internal/auth"
	"go-slotcars/internal/db"
	"go-slotcars/internal/router"
	"go-slotcars/internal/users"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Setup the database connection
	dbConn, err := db.SetupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories and use cases
	userRepo := users.NewSQLiteUserRepository(dbConn)
	userUC := users.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := users.NewUserHandler(userUC)

	// Initialize JWT Manager and middleware
	jwtManager := auth.NewJWTManager("supersecretkey")
	authMiddleware := auth.NewAuthMiddleware(jwtManager)

	// Setup Fiber app and routes
	app := fiber.New()
	router.SetupRoutes(app, userHandler, authMiddleware)

	log.Println("🚀 Server is running at http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
