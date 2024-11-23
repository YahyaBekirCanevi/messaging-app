package model

import "time"

// Message represents a message document in MongoDB
type Message struct {
	ID        string    `bson:"id,omitempty" json:"id"`       // Custom ID field
	Sender    string    `bson:"sender" json:"sender"`         // Sender of the message
	Receiver  string    `bson:"receiver" json:"receiver"`     // Receiver of the message
	Content   string    `bson:"content" json:"content"`       // Message content
	CreatedAt time.Time `bson:"created_at" json:"created_at"` // Timestamp for when the message was created
}
