package config

import (
	"golang-app-todolist/internal/delivery/rest/handler"
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

	// service
	userService := service.UserServiceImpl(userRepository, config.Validation, config.Log, config.Viper)

	// handler
	userHandler := handler.USerHandlerImpl(userService)

	// route
	route := &routes.RouteConfig{
		App:         config.App,
		UserHandler: userHandler,
	}
	route.Setup()
}
