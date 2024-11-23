package routes

import (
	"messaging-app/controller"

	"github.com/gorilla/mux"
)

// RegisterMessageRoutes defines routes for the messaging app
func RegisterMessageRoutes(router *mux.Router, controller *controller.MessageController) {
	router.HandleFunc("/messages", controller.GetAllMessages).Methods("GET")
	router.HandleFunc("/messages", controller.CreateMessage).Methods("POST")
	router.HandleFunc("/messages/{id}", controller.GetMessageByID).Methods("GET")
	router.HandleFunc("/messages/{id}", controller.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/messages/user/{user}", controller.GetUserMessages).Methods("GET")
}
