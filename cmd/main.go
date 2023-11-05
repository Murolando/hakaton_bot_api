package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Murolando/hakaton_bot_api/pkg/handler"
	"github.com/joho/godotenv"

	srv "github.com/Murolando/hakaton_bot_api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env variables:", err)
	}
	serverPort := os.Getenv("SERVER_PORT")

	handler := handler.NewHandler()
	s := new(srv.Server)
	if err := s.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	fmt.Println('o')
}
