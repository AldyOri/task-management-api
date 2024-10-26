package dto

import (
	"time"
	dtoImage "todo-app/models/dto/dto-image"
)

type TaskResponse struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Images      []dtoImage.ImageResponse `json:"images"`
	Completed   bool           `json:"completed"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
