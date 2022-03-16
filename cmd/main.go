package main

import (
	"log"
	"restAPI"
	"restAPI/pkg/handler"
	"restAPI/pkg/repository"
	"restAPI/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restAPI.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server, %s", err.Error())
	}
}
