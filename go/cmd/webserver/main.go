package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/infra"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/interfaces/webserver"
	"github.com/marcosbitetti/godot-manage-attribuitions-plugin/intrenal/usecases"
)

func main() {

	// get injected database_path: ex: DATABASE_PATH=./database.db ./bin/attribuitions-local-server-amd64-linux
	path := os.Getenv("DATABASE_PATH")

	storage, err := infra.NewStorage(path)
	if err != nil {
		println(string(usecases.FormatJSON(nil, err)))
		return
	}
	defer storage.CloseDatabase()

	errorChan := make(chan error)

	server := webserver.NewHttpServer(storage, errorChan)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
	case err := <-errorChan:
		log.Println(err.Error())
	}
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
