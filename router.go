package main

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/goodbye", GoodbyeHandler)
}
