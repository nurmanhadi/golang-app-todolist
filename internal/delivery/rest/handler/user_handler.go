package handler

import (
	"golang-app-todolist/internal/model"
	"golang-app-todolist/internal/service"
	"golang-app-todolist/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type USerHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
type userHandler struct {
	userService service.UserService
}

func USerHandlerImpl(userService service.UserService) USerHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.userService.Register(*request); err != nil {
		return err
	}
	return ResponseHandler(c, 200, "OK")
}
func (h *userHandler) Login(c *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	if err := c.BodyParser(request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	token, err := h.userService.Login(*request)
	if err != nil {
		return err
	}
	return ResponseHandler(c, 200, token)
}
