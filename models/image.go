package models

import "time"

type Image struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	TaskID      uint      `json:"task_id" gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Filename    string    `json:"filename" gorm:"not null"`
	Data        []byte    `json:"data" gorm:"bytea;not null"`
	ContentType string    `json:"content_type" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
