package model

import "time"

type GroupsUsers struct {
	AbstractModel
	UserIDJoin         string     `json:"user_join"`
	LastDeletedMessage *time.Time `json:"last_deleted_messages"`
}
