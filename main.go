package main

import (
	"log"
)

func main() {
	Connect("")

	r := SetupRouter()

	log.Println("Server running on http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
