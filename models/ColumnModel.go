package models

type ColumnModel struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" validate:"required,min=1,max=255"`
	Position  uint   `json:"position"`
	ProjectId int64  `json:"project_id"`
}
