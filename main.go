package main

import (
	"log"
	"net/http"

	"messaging-app/controller"
	"messaging-app/database"
	"messaging-app/repository"
	"messaging-app/routes"
	"messaging-app/service"

	"github.com/gorilla/mux"
)

func main() { // Initialize MongoDB connection
	database.ConnectDB()
	defer database.DisconnectDB()

	// Repositories and Services
	messageRepo := repository.NewMessageRepository()
	messageService := service.NewMessageService(messageRepo)

	// Initialize controllers with the service layer
	messageController := controller.NewMessageController(messageService)

	// Initialize router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterMessageRoutes(router, messageController)

	// Start server
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
