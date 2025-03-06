package main

import (
	"fmt"
	"frontend3_server2/internal"
	"log"
	"net/http"
	"sync"
)

// @title MEOW
// @version 1.0
// @BasePath /
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		db, err := internal.InitDataBase()
		if err != nil {
			log.Fatal(err)
		}
		handler := &internal.Handler{db}
		handler.InitRoutes()
	}()
	go func() {
		defer wg.Done()
		http.HandleFunc("/ws", internal.HandleConnections)

		fmt.Println("WebSocket server started on :10003")
		err := http.ListenAndServe(":10003", nil)
		if err != nil {
			fmt.Println("ListenAndServe:", err)
		}
	}()
	wg.Wait()
}
