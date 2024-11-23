package service

import (
	"context"
	"errors"
	"messaging-app/model"
	"messaging-app/repository"
	"strings"
	"time"

	"github.com/google/uuid"
)

// MessageService provides methods to handle business logic for messages
type MessageService struct {
	repo *repository.MessageRepository
}

// NewMessageService creates a new instance of MessageService
func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

// CreateMessage adds business logic and saves the message using the repository
func (ms *MessageService) CreateMessage(ctx context.Context, message *model.Message) (*model.Message, error) {
	// Generate a unique, lowercase ID for the message
	message.ID = strings.ToLower(uuid.New().String())
	message.CreatedAt = time.Now()

	// Validate the message fields (add your validation logic here)
	if message.Sender == "" || message.Receiver == "" || message.Content == "" {
		return nil, errors.New("userId and content are required")
	}

	err := ms.repo.Save(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// GetAllMessages retrieves all messages using the repository
func (ms *MessageService) GetAllMessages(ctx context.Context) ([]model.Message, error) {
	return ms.repo.FindAll(ctx)
}

// GetMessageByID retrieves a specific message by ID using the repository
func (ms *MessageService) GetMessageByID(ctx context.Context, id string) (*model.Message, error) {
	return ms.repo.FindByID(ctx, strings.ToLower(id))
}

// DeleteMessage deletes a message by ID using the repository
func (ms *MessageService) DeleteMessage(ctx context.Context, id string) error {
	return ms.repo.DeleteByID(ctx, strings.ToLower(id))
}

// GetMessagesForUser retrieves messages where the user is sender or receiver
func (ms *MessageService) GetMessagesForUser(ctx context.Context, user string) ([]model.Message, error) {
	return ms.repo.FindBySenderOrReceiver(ctx, user)
}
