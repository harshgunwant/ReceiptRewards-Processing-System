package main

import (
	"FetchRewardsAssessment/internal/routes"
	"log"
)

func main() {
	// Initialize Gin router
	router := routes.RegisterRoutes()

	// Start the server
	log.Println("Welcome to fetch: Starting server on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
