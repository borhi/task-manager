package models

type ColumnModel struct {
	Id        uint   `json:"id"`
	Name      string `json:"name" validate:"required,unique,min=1,max=255"`
	Position  uint   `json:"position"`
	ProjectId uint   `json:"project_id"`
}
