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
	ItemHandler      handler.ItemHandler
}

func (r *RouteConfig) Setup() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", r.UserHandler.Register)
	auth.Post("/login", r.UserHandler.Login)

	checklist := api.Group("/checklist", r.AuthMiddleware.Auth())
	checklist.Post("/", r.ChecklistHandler.Add)
	checklist.Get("/", r.ChecklistHandler.FindAll)
	checklist.Get("/:checklistId", r.ChecklistHandler.FindById)
	checklist.Delete("/:checklistId", r.ChecklistHandler.Delete)

	item := checklist.Group("/:checklistId/item")
	item.Post("/", r.ItemHandler.Add)
	item.Get("/:itemId", r.ItemHandler.FindById)
	item.Put("/:itemId", r.ItemHandler.UpdateStatus)
	item.Put("/rename/:itemId", r.ItemHandler.UpdateItemName)
	item.Delete("/:itemId", r.ItemHandler.Delete)

}
