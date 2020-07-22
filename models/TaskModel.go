package models

type TaskModel struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"min=0,max=5000"`
	Position    int    `json:"position"`
	ColumnId    int64  `json:"column_id"`
}
