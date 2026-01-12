package main

/*import (
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
}*/

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

var messageService = NewMessageService()

func HelloHandler(c *gin.Context) {
	c.JSON(200, Response{
		Message: messageService.HelloMessage(),
	})
}

func GoodbyeHandler(c *gin.Context) {
	c.JSON(200, Response{
		Message: messageService.GoodbyeMessage(),
	})
}
