package repository

import "gitlab.com/vdat/mcsvc/chat/pkg/model"

type GroupRepo interface {
	GetListGroupByUser(subUser string) ([]model.Groups, error)
}
