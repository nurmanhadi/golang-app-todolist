package entity

type Checklist struct {
	ID           int    `gorm:"primaryKey;auto_increment" json:"id"`
	UserUsername string `json:"user_username"`
	Name         string `json:"name"`
}
