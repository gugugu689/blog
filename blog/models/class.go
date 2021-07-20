package models

import "time"

type Class struct {
	ID int64 `json:"id" db:"class_id"`
	Name string `json:"name" db:"class_name" binding:"required"`
}

type ClassDetail struct {
	ID int64 `json:"id" db:"class_id"`
	Name string `json:"name" db:"class_name"`
	CreateTime time.Time `json:"creat_time" db:"create_time"`
}