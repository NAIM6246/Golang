package main

import (
	"Golang/config"
	"Golang/conn"
	"Golang/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":3000"
	rout := chi.NewRouter()
	userHandler := handler.NewUserHandler()
	rout.Route("/users", userHandler.Handle)
	fmt.Println("Serving on port ", port)

	http.ListenAndServe(port, rout)
	config2 := config.NewDBConfig()
	connection := conn.ConnectDB(config2)
	defer connection.Close()
}
