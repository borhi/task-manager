package models

import "time"

type CommentModel struct {
	Id        uint      `json:"id"`
	Text      string    `json:"text" validate:"required,unique,min=1,max=5000"`
	CreatedAt time.Time `json:"created_at"`
	TaskId    uint      `json:"task_id"`
}
