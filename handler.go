package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

var messageService = NewMessageService()

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: messageService.HelloMessage(),
	})
}

func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: messageService.GoodbyeMessage(),
	})
}
