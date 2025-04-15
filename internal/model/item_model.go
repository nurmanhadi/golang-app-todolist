package model

type ItemAddRequest struct {
	ItemName    string `json:"item_name" validate:"required,max=100"`
	Description string `json:"description,omitempty"`
}
type ItemUpdateRequest struct {
	ItemName string `json:"item_name" validate:"required,max=100"`
}
