package model

import "time"

type AbstractModel struct{
	ID        uint `json:"id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}