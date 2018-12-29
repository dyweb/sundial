package message

import (
	"context"

	"github.com/dyweb/sundial/pkg/models"
	"github.com/dyweb/sundial/pkg/store/rdb"
)

// ListMessages returns all messages.
func ListMessages(ctx context.Context, count int) ([]models.Message, error) {
	ds := rdb.FromContext(ctx)
	return ds.GetMessages()
}

// GetMessage return a message by id.
func GetMessage(ctx context.Context, id int) (*models.Message, error) {
	ds := rdb.FromContext(ctx)
	return ds.GetMessage(id)
}
