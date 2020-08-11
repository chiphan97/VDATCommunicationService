package repository

import "golangproject/model"

type ChatBoxRepo interface {
	GetChatBoxs() ([]model.ChatBoxModel, error)
}