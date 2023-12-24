package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kedarnathpc/go-postgres/router"
)

func main() {
	// Initialize the router
	r := router.Router()

	fmt.Println("Starting server at port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
