package repository

import (
	"golang-app-todolist/internal/entity"

	"gorm.io/gorm"
)

type ChecklistRepository interface {
	Save(checklist entity.Checklist) error
	FindAll(userUsername string) ([]entity.Checklist, error)
	Delete(checklistId int) error
	Count(checklistId int) (int64, error)
}
type checklistRepository struct {
	db *gorm.DB
}

func ChecklistRepositoryImpl(db *gorm.DB) ChecklistRepository {
	return &checklistRepository{db: db}
}
func (r *checklistRepository) Save(checklist entity.Checklist) error {
	return r.db.Save(&checklist).Error
}
func (r *checklistRepository) FindAll(userUsername string) ([]entity.Checklist, error) {
	var checklists []entity.Checklist
	err := r.db.Where("user_username = ?", userUsername).Find(&checklists).Error
	if err != nil {
		return nil, err
	}
	return checklists, err
}
func (r *checklistRepository) Delete(checklistId int) error {
	checklist := new(entity.Checklist)
	err := r.db.Where("id = ?", checklistId).Delete(&checklist).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *checklistRepository) Count(checklistId int) (int64, error) {
	var count int64
	checklist := new(entity.Checklist)
	err := r.db.Model(&checklist).Where("id = ?", checklistId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
