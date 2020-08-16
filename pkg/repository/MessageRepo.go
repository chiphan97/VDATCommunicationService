package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type MessageRepo interface {
	GetMessages() ([]model.MessageModel, error)
}
