package main

import (
	"log"

	"github.com/Orazali123/todo-app"
	"github.com/Orazali123/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRountes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
