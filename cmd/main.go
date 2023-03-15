package main

import (
	todo "github.com/antonchaban/file-hub-go"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal("error occurred while running http server: ", err.Error())
	}
}
