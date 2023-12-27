package main

import (
	"log"

	"github.com/bostud/blacklist"
	"github.com/bostud/blacklist/pkg/handler"
	"github.com/bostud/blacklist/pkg/repository"
	"github.com/bostud/blacklist/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Config reading error: %s", err.Error())
	}
	repos := repository.Newrepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	server := new(blacklist.Server)
	port := viper.GetString("port")

	if err := server.Run(port, handler.InitRoutes()); err != nil {
		log.Fatalf("Error during server run: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
