package datastore

import (
	"github.com/dyweb/sundial/pkg/models"
)

func (ds *datastore) GetMessages() ([]models.Message, error) {
	messages := []models.Message{}
	err := ds.Find(messages).Error
	return messages, err
}

func (ds *datastore) GetMessage(id int) (*models.Message, error) {
	message := &models.Message{}
	err := ds.First(message, "id = ?", id).Error
	return message, err
}
