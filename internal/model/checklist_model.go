package model

type ChecklistAddRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}
