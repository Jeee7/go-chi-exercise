package main

import (
	"go-fcm-app/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("Server running at :8080")
	log.Println(http.ListenAndServe(":8080", r))
}
