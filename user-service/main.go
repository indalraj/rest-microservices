package main

import (
	"fmt"
	"net/http"

	"user-service/handler"
)

func main() {
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("User Service running on :8081")
	http.ListenAndServe(":8081", nil)
}
