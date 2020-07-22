package models

type ProjectModel struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"min=0,max=1000"`
}
