package handler

import (
	"golang-app-todolist/internal/model"
	"golang-app-todolist/internal/service"
	"golang-app-todolist/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler interface {
	Add(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	UpdateItemName(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
type itemHandler struct {
	itemService service.ItemService
}

func ItemHandlerImpl(itemService service.ItemService) ItemHandler {
	return &itemHandler{
		itemService: itemService,
	}
}
func (h *itemHandler) Add(c *fiber.Ctx) error {
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	checklistId := c.Params("checklistId")
	request := new(model.ItemAddRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.itemService.Add(checklistId, *request); err != nil {
		return err
	}
	return ResponseHandler(c, 201, "created")
}
func (h *itemHandler) FindById(c *fiber.Ctx) error {
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	checklistId := c.Params("checklistId")
	itemId := c.Params("itemId")
	item, err := h.itemService.FindById(checklistId, itemId)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, item)
}
func (h *itemHandler) UpdateStatus(c *fiber.Ctx) error {
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	checklistId := c.Params("checklistId")
	itemId := c.Params("itemId")
	err := h.itemService.UpdateStatus(checklistId, itemId)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, "OK")
}
func (h *itemHandler) UpdateItemName(c *fiber.Ctx) error {
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	checklistId := c.Params("checklistId")
	itemId := c.Params("itemId")
	request := new(model.ItemUpdateRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	err := h.itemService.UpdateItemName(checklistId, itemId, *request)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, "OK")
}
func (h *itemHandler) Delete(c *fiber.Ctx) error {
	_, ok := c.Locals("username").(string)
	if !ok {
		return exception.NewError(401, "unauthorized")
	}
	checklistId := c.Params("checklistId")
	itemId := c.Params("itemId")
	err := h.itemService.Delete(checklistId, itemId)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, "OK")
}
