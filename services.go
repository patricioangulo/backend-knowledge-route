package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MessageService struct{}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (s *MessageService) HelloMessage() string {
	return "Hello!"
}

func (s *MessageService) GoodbyeMessage() string {
	return "Goodbye!"
}

func CreateBooking(b Booking) (*mongo.InsertOneResult, error) {
	collection := GetCollection("bookings")

	// Assign a new ObjectID
	b.ID = bson.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
