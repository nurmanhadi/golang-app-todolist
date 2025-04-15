package entity

type Item struct {
	ID          int    `gorm:"primaryKey,auto_increment" json:"id"`
	ChecklistId int    `json:"checklist_id"`
	ItemName    string `json:"item_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
