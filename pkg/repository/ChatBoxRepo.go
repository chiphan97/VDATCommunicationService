package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type ChatBoxRepo interface {
	GetChatBoxs() ([]model.ChatBoxModel, error)
}
