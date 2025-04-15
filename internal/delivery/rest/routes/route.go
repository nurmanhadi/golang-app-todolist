package routes

import (
	"golang-app-todolist/internal/delivery/rest/handler"
	"golang-app-todolist/internal/delivery/rest/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App              *fiber.App
	AuthMiddleware   *middleware.MiddlewareConfig
	UserHandler      handler.USerHandler
	ChecklistHandler handler.ChecklistHandler
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", r.UserHandler.Register)
	auth.Post("/login", r.UserHandler.Login)

	checklist := api.Group("/checklist", r.AuthMiddleware.Auth())
	checklist.Post("/", r.ChecklistHandler.Add)
	checklist.Get("/", r.ChecklistHandler.FindAll)
	checklist.Delete("/:checklistId", r.ChecklistHandler.Delete)
}
