package main

import (
	"github.com/spf13/viper"
	"log"
	"server"
	"server/package/handler"
	"server/package/repository"
	"server/package/service"
)

func main() {
	//handlers := new(handler.Handler)
	if err := initConfig(); err != nil {
		log.Fatalf("error iniitializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}