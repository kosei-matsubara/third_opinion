package main

import (
	"log"

	"github.com/third_opinion/backend/internal/interfaces/router"
)

func main() {
	// Setup router with all Todo API endpoints
	r := router.SetupRouter()

	// Start server
	log.Println("Starting Todo API server on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
