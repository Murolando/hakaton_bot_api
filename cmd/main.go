package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Murolando/hakaton_bot_api/pkg/handler"

	srv "github.com/Murolando/hakaton_bot_api"
)

func main() {

	serverPort := os.Getenv("SERVER_PORT")

	handler := handler.NewHandler()

	s := new(srv.Server)
	if err := s.Run(serverPort, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	fmt.Println('o')
}
