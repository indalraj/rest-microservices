package main

import (
	"fmt"
	"log"
	"net/http"

	"user-service/config"
	"user-service/handler"
)

func main() {
	config.ConnectDB()
	config.Migrate()

	http.HandleFunc("/users", handler.UsersHandler)

	fmt.Println("User Service running on :8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
