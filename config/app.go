package config

import (
	"golang-app-todolist/internal/delivery/rest/handler"
	"golang-app-todolist/internal/delivery/rest/middleware"
	"golang-app-todolist/internal/delivery/rest/routes"
	"golang-app-todolist/internal/repository"
	"golang-app-todolist/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Viper      *viper.Viper
	DB         *gorm.DB
	App        *fiber.App
	Validation *validator.Validate
	Log        *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	// repository
	userRepository := repository.UserRepositoryImpl(config.DB)
	checklistRepository := repository.ChecklistRepositoryImpl(config.DB)

	// service
	userService := service.UserServiceImpl(userRepository, config.Validation, config.Log, config.Viper)
	checklistService := service.ChecklistServiceImpl(checklistRepository, config.Validation, config.Log)

	// handler
	userHandler := handler.USerHandlerImpl(userService)
	checklistHandler := handler.ChecklistHandlerImpl(checklistService)

	// middleware
	middleware := &middleware.MiddlewareConfig{
		Viper: config.Viper,
		Log:   config.Log,
	}

	// route
	route := &routes.RouteConfig{
		App:              config.App,
		AuthMiddleware:   middleware,
		UserHandler:      userHandler,
		ChecklistHandler: checklistHandler,
	}
	route.Setup()
}
