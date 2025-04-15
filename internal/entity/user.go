package entity

type User struct {
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"primaryKey" json:"username"`
	Password string `json:"password"`
}
