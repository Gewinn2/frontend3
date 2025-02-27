package main

import (
	"log"
)

func main() {
	db, err := InitDataBase()
	if err != nil {
		log.Fatal(err)
	}
	handler := &Handler{db}
	handler.InitRoutes()
}
