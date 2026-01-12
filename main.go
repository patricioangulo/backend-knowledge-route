package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// RegisterRoutes()
	r := SetupRouter()
	r.Run(":8080")

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
