package repository

import (
	"golang-app-todolist/internal/entity"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Save(item entity.Item) error
	FindById(checklistId int, itemId int) (*entity.Item, error)
	UpdateStatus(checklistId int, itemId int, status string) error
	UpdateItemName(checklistId int, itemId int, itemName string) error
	Count(itemId int) (int64, error)
	Delete(itemId int) error
}
type itemRepository struct {
	db *gorm.DB
}

func ItemRepositoryImpl(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Save(item entity.Item) error {
	return r.db.Save(&item).Error
}
func (r *itemRepository) FindById(checklistId int, itemId int) (*entity.Item, error) {
	item := new(entity.Item)
	err := r.db.Where("id = ? AND checklist_id = ?", itemId, checklistId).First(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}
func (r *itemRepository) UpdateStatus(checklistId int, itemId int, status string) error {
	item := new(entity.Item)
	err := r.db.Model(&item).Where("id = ? AND checklist_id = ?", itemId, checklistId).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *itemRepository) UpdateItemName(checklistId int, itemId int, itemName string) error {
	item := new(entity.Item)
	err := r.db.Model(&item).Where("id = ? AND checklist_id = ?", itemId, checklistId).Update("item_name", itemName).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *itemRepository) Count(itemId int) (int64, error) {
	item := new(entity.Item)
	var count int64
	err := r.db.Model(&item).Where("id = ?", itemId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *itemRepository) Delete(itemId int) error {
	item := new(entity.Item)
	err := r.db.Where("id = ?", itemId).Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
