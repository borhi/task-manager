package models

type TaskModel struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" validate:"required,min=1,max=500"`
	Description string `json:"description" validate:"min=0,max=5000"`
	Position    uint   `json:"position"`
	ColumnId    uint   `json:"column_id"`
}
