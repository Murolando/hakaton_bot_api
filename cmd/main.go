package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Murolando/default_rest_arch/pkg/handler"
	"github.com/Murolando/default_rest_arch/pkg/repository"
	"github.com/Murolando/default_rest_arch/pkg/repository/postgres"
	"github.com/Murolando/default_rest_arch/pkg/service"

	srv "github.com/Murolando/default_rest_arch"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPas := os.Getenv("DB_PASSWORD")
	serverPort := os.Getenv("SERVER_PORT")
	dbConfig := postgres.NewConfig(dbHost, dbPort, dbUserName, dbPas, dbName)
	db, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	s := new(srv.Server)
	if err := s.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	fmt.Println('o')
}
