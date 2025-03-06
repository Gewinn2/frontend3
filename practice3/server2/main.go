package main

import "log"

// @title MEOW
// @version 1.0
// @BasePath /
func main() {
	db, err := InitDataBase()
	if err != nil {
		log.Fatal(err)
	}
	handler := &Handler{db}
	handler.InitRoutes()
}
