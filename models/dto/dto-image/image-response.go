package dtoImage

import "time"

type ImageResponse struct {
	ID          uint      `json:"id"`
	Filename    string    `json:"filename"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
}
