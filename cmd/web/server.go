package main

import (
	"log"
	"net/http"

	"dach-trier.com/web"
)

func main() {
	log.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", web.NewApp().Router())
	log.Fatal(err)
}
