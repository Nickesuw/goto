package main

import (
	todo "goto-app"
	"goto-app/pkg/handler"
	"log"
)

func main() {
	handlers:= new(handler.Handler)
	serv:=new(todo.Server)
	if err:= serv.Run("8000",handlers.InitRoutes()); err!=nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
