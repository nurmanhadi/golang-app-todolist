package handler

import (
	"golang-app-todolist/internal/model"
	"golang-app-todolist/internal/service"
	"golang-app-todolist/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type ChecklistHandler interface {
	Add(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}
type checklistHandler struct {
	checklistService service.ChecklistService
}

func ChecklistHandlerImpl(checklistService service.ChecklistService) ChecklistHandler {
	return &checklistHandler{
		checklistService: checklistService,
	}
}
func (h *checklistHandler) Add(c *fiber.Ctx) error {
	username, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	request := new(model.ChecklistAddRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.checklistService.Add(username, *request); err != nil {
		return err
	}
	return ResponseHandler(c, 201, "Created")
}
func (h *checklistHandler) FindAll(c *fiber.Ctx) error {
	username, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}

	checklists, err := h.checklistService.FindAll(username)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 201, checklists)
}
func (h *checklistHandler) Delete(c *fiber.Ctx) error {
	checklistId := c.Params("checklistId")
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}

	err := h.checklistService.Delete(checklistId)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, "OK")
}
func (h *checklistHandler) FindById(c *fiber.Ctx) error {
	checklistId := c.Params("checklistId")
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}

	checklist, err := h.checklistService.FindById(checklistId)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, checklist)
}
