package main

import (
	"fmt"
	"golang-app-todolist/config"
)

func main() {
	viper := config.NewViper()
	db := config.NewConnection(viper)
	log := config.NewLogger()
	validation := config.NewValidator()
	app := config.NewFiber(viper)
	config.Bootstrap(&config.BootstrapConfig{
		Viper:      viper,
		DB:         db,
		App:        app,
		Validation: validation,
		Log:        log,
	})
	port := viper.GetInt("server.port")
	host := viper.GetString("server.host")
	if err := app.Listen(fmt.Sprintf("%s:%d", host, port)); err != nil {
		panic(err)
	}
}
