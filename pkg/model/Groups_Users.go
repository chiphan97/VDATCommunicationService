package model

import "time"

type GroupsUsers struct {
	AbstractModel
	SubUserJoin        string     `json:"sub_user_join"`
	LastDeletedMessage *time.Time `json:"last_deleted_messages"`
}
