package models

import "time"

type CommentModel struct {
	Id        int64     `json:"id"`
	Text      string    `json:"text" validate:"required,min=1,max=5000"`
	CreatedAt time.Time `json:"created_at"`
	TaskId    int64     `json:"task_id"`
}
