package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"dach-trier.com/web"
)

func main() {
	app, err := web.NewApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to initialize web application\n")
		fmt.Fprintf(os.Stderr, "reason: %v\n", err)

		os.Exit(1)
	}
	log.Println("Server is running at http://localhost:8080")
	err = http.ListenAndServe(":8080", app.Router())
	log.Fatal(err)
}
