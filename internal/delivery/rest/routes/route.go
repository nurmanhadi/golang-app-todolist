package routes

import (
	"golang-app-todolist/internal/delivery/rest/handler"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App         *fiber.App
	UserHandler handler.USerHandler
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", r.UserHandler.Register)
	auth.Post("/login", r.UserHandler.Login)
}
