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
	"net/http"

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

func CreateBookingHandler(c *gin.Context) {
	var booking Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created",
		"id":      result.InsertedID,
	})
}
