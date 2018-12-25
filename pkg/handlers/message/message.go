package message

import (
	"context"
	"fmt"
)

// Message describes a message entry.
type Message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ListMessages returns all messages.
func ListMessages(ctx context.Context, count int) ([]Message, error) {
	messages := make([]Message, count)
	for i := 0; i < count; i++ {
		messages[i].ID = i
		messages[i].Title = fmt.Sprintf("Example %d", i)
		messages[i].Content = fmt.Sprintf("Content of example %d", i)
	}
	return messages, nil
}

// GetMessage return a message by id.
func GetMessage(ctx context.Context, id int) (*Message, error) {
	return &Message{
		ID:      id,
		Title:   "This is an example",
		Content: "Example content",
	}, nil
}
