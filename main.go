package main

import (
	"github.com/pete911/rides/pkg/mongo"
	"github.com/pete911/rides/web/app"
	"log"
)

func main() {

	mongo.InitSession("localhost:27017")
	server := app.NewRidesServer(8080)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
