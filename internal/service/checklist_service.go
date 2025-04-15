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

type ChecklistService interface {
	Add(userUsername string, request model.ChecklistAddRequest) error
	FindAll(userUsername string) ([]entity.Checklist, error)
	Delete(checklistId string) error
	FindById(checklistId string) (*entity.Checklist, error)
}
type checklistService struct {
	checklistRepository repository.ChecklistRepository
	validation          *validator.Validate
	log                 *logrus.Logger
}

func ChecklistServiceImpl(
	checklistRepository repository.ChecklistRepository,
	validation *validator.Validate,
	log *logrus.Logger,
) ChecklistService {
	return &checklistService{
		checklistRepository: checklistRepository,
		validation:          validation,
		log:                 log,
	}
}

func (s *checklistService) Add(userUsername string, request model.ChecklistAddRequest) error {
	if err := s.validation.Struct(&request); err != nil {
		s.log.WithError(err).Warn("validation error")
		return err
	}
	checklist := &entity.Checklist{
		UserUsername: userUsername,
		Name:         request.Name,
	}
	if err := s.checklistRepository.Save(*checklist); err != nil {
		s.log.WithError(err).Error("failed add checklist to database")
		return err
	}
	return nil
}
func (s *checklistService) FindAll(userUsername string) ([]entity.Checklist, error) {
	checklists, err := s.checklistRepository.FindAll(userUsername)
	if err != nil {
		s.log.WithError(err).Error("failed find checklists to database")
		return nil, err
	}
	return checklists, nil
}
func (s *checklistService) Delete(checklistId string) error {
	id, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	count, err := s.checklistRepository.Count(id)
	if err != nil {
		s.log.WithError(err).Error("failed count checklist to database")
		return err
	}

	if count < 1 {
		s.log.WithError(err).Warn("checklist not found")
		return exception.NewError(404, "checklist not found")
	}
	err = s.checklistRepository.Delete(id)
	if err != nil {
		s.log.WithError(err).Error("failed delete checklist to database")
		return err
	}
	return nil
}
func (s *checklistService) FindById(checklistId string) (*entity.Checklist, error) {
	newChecklistId, err := strconv.Atoi(checklistId)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	checklist, err := s.checklistRepository.FindById(newChecklistId)
	if err != nil {
		s.log.WithError(err).Error("failed get checklist by id to dayabase")
		return nil, err
	}
	return checklist, nil
}
