package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"messaging-app/model"
	"messaging-app/service"

	"github.com/gorilla/mux"
)

// MessageService provides methods to handle business logic for messages
type MessageController struct {
	service *service.MessageService
}

// NewMessageService creates a new instance of MessageService
func NewMessageController(service *service.MessageService) *MessageController {
	return &MessageController{
		service: service,
	}
}

// CreateMessage handles creating a new message
func (controller *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := controller.service.CreateMessage(context.TODO(), &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

// GetAllMessages handles fetching all messages
func (controller *MessageController) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := controller.service.GetAllMessages(context.TODO())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// GetMessageByID handles fetching a message by ID
func (controller *MessageController) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	message, err := controller.service.GetMessageByID(context.TODO(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// DeleteMessage handles deleting a message by ID
func (controller *MessageController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := controller.service.DeleteMessage(context.TODO(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUserMessages retrieves all messages for a user as sender or receiver
func (controller *MessageController) GetUserMessages(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	messages, err := controller.service.GetMessagesForUser(context.TODO(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
