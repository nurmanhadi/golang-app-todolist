package repository

import (
	"golang-app-todolist/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) error
	FindByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) Save(user entity.User) error {
	return r.db.Save(&user).Error
}
func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
