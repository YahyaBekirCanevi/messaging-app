package repository

import (
	"context"
	"errors"
	"messaging-app/database"
	"messaging-app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MessageRepository handles raw MongoDB operations
type MessageRepository struct{
	
}

// NewMessageRepository creates a new instance of MessageRepository
func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

// Save inserts a new message into MongoDB
func (repo *MessageRepository) Save(ctx context.Context, message *model.Message) error {
	message.CreatedAt = time.Now()

	_, err := database.MessageCollection.InsertOne(ctx, message)
	if err != nil {
		return errors.New("failed to save message: " + err.Error())
	}

	return nil
}

// FindAll retrieves all messages
func (repo *MessageRepository) FindAll(ctx context.Context) ([]model.Message, error) {
	cursor, err := database.MessageCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.New("failed to retrieve messages: " + err.Error())
	}
	defer cursor.Close(ctx)

	var messages []model.Message
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, errors.New("failed to decode messages: " + err.Error())
	}

	return messages, nil
}

// FindByID retrieves a message by ID
func (repo *MessageRepository) FindByID(ctx context.Context, id string) (*model.Message, error) {
	var message model.Message
	err := database.MessageCollection.FindOne(ctx, bson.M{"id": id}).Decode(&message)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("message not found")
	}
	if err != nil {
		return nil, errors.New("failed to retrieve message: " + err.Error())
	}

	return &message, nil
}

// FindBySenderOrReceiver retrieves messages where the user is sender or receiver
func (repo *MessageRepository) FindBySenderOrReceiver(ctx context.Context, user string) ([]model.Message, error) {
	cursor, err := database.MessageCollection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"sender": user},
			{"receiver": user},
		},
	})
	if err != nil {
		return nil, errors.New("failed to retrieve messages: " + err.Error())
	}
	defer cursor.Close(ctx)

	var messages []model.Message
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, errors.New("failed to decode messages: " + err.Error())
	}

	return messages, nil
}

// DeleteByID deletes a message by ID
func (repo *MessageRepository) DeleteByID(ctx context.Context, id string) error {
	_, err := database.MessageCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return errors.New("failed to delete message: " + err.Error())
	}

	return nil
}
