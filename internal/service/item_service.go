package service

import (
	"golang-app-todolist/internal/entity"
	"golang-app-todolist/internal/model"
	"golang-app-todolist/internal/repository"
	"golang-app-todolist/pkg/exception"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ItemService interface {
	Add(checklistId string, request model.ItemAddRequest) error
	FindById(checklistId string, itemId string) (*entity.Item, error)
	UpdateStatus(checklistId string, itemId string) error
	UpdateItemName(checklistId string, itemId string, request model.ItemUpdateRequest) error
	Delete(checklistId string, itemId string) error
}
type itemService struct {
	itemRepository      repository.ItemRepository
	checklistRepository repository.ChecklistRepository
	validation          *validator.Validate
	log                 *logrus.Logger
}

func ItemServiceImpl(
	itemRepository repository.ItemRepository,
	checklistRepository repository.ChecklistRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) ItemService {
	return &itemService{
		itemRepository:      itemRepository,
		checklistRepository: checklistRepository,
		validation:          validation,
		log:                 log,
	}
}

func (s *itemService) Add(checklistId string, request model.ItemAddRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.log.WithError(err).Warn("validation error")
		return err
	}
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countChecklist, err := s.checklistRepository.Count(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return err
	}

	if countChecklist < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return exception.NewError(404, "checklist not found")
	}

	item := &entity.Item{
		ChecklistId: newChecklistId,
		ItemName:    request.ItemName,
		Description: request.Description,
	}
	err = s.itemRepository.Save(*item)
	if err != nil {
		s.log.WithError(err).Error("failed save item to database")
		return err
	}
	return nil
}
func (s *itemService) FindById(checklistId string, itemId string) (*entity.Item, error) {
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	countChecklist, err := s.checklistRepository.Count(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return nil, err
	}

	if countChecklist < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return nil, exception.NewError(404, "checklist not found")
	}
	newItemId, err := strconv.Atoi(itemId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	item, err := s.itemRepository.FindById(newChecklistId, newItemId)
	if err != nil {
		s.log.WithError(err).Warn("failed get data from database")
		return nil, exception.NewError(404, "item not found")
	}
	return item, nil
}
func (s *itemService) UpdateStatus(checklistId string, itemId string) error {
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countChecklist, err := s.checklistRepository.Count(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return err
	}

	if countChecklist < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return exception.NewError(404, "checklist not found")
	}
	newItemId, err := strconv.Atoi(itemId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	item, err := s.itemRepository.FindById(newChecklistId, newItemId)
	if err != nil {
		s.log.WithError(err).Warn("failed get data from database")
		return exception.NewError(404, "item not found")
	}
	if item.Status == "pending" {
		status := "completed"
		err := s.itemRepository.UpdateStatus(newChecklistId, newItemId, status)
		if err != nil {
			s.log.WithError(err).Error("failed update item to database")
			return err
		}
	}
	return nil
}
func (s *itemService) UpdateItemName(checklistId string, itemId string, request model.ItemUpdateRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.log.WithError(err).Warn("validation error")
		return err
	}
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countChecklist, err := s.checklistRepository.Count(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return err
	}

	if countChecklist < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return exception.NewError(404, "checklist not found")
	}
	newItemId, err := strconv.Atoi(itemId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countItem, err := s.itemRepository.Count(newItemId)
	if err != nil {
		s.log.WithError(err).Error("failed count item to database")
		return err
	}
	if countItem < 1 {
		s.log.WithError(err).Warn("item not found")
		return exception.NewError(404, "item not found")
	}
	err = s.itemRepository.UpdateItemName(newChecklistId, newItemId, request.ItemName)
	if err != nil {
		s.log.WithError(err).Error("failed update item name to database")
		return err
	}
	return nil
}
func (s *itemService) Delete(checklistId string, itemId string) error {
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countChecklist, err := s.checklistRepository.Count(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return err
	}

	if countChecklist < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return exception.NewError(404, "checklist not found")
	}
	newItemId, err := strconv.Atoi(itemId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countItem, err := s.itemRepository.Count(newItemId)
	if err != nil {
		s.log.WithError(err).Error("failed count item to database")
		return err
	}
	if countItem < 1 {
		s.log.WithError(err).Warn("item not found")
		return exception.NewError(404, "item not found")
	}
	err = s.itemRepository.Delete(newItemId)
	if err != nil {
		s.log.WithError(err).Error("failed delete item to database")
		return err
	}
	return nil
}
