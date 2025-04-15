package repository

import (
	"golang-app-todolist/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) error
	FindByUsername(username string) (*entity.User, error)
	CountByUsername(username string) (int64, error)
	CountByEmail(email string) (int64, error)
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
	return user, err
}
func (r *userRepository) CountByUsername(username string) (int64, error) {
	user := new(entity.User)
	var count int64
	err := r.db.Model(user).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
func (r *userRepository) CountByEmail(email string) (int64, error) {
	user := new(entity.User)
	var count int64
	err := r.db.Model(user).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
