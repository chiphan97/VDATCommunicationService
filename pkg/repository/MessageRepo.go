package repository

import "golangproject/model"

type MessageRepo interface {
	GetMessages() ([]model.MessageModel, error)
}