package main

import "go.mongodb.org/mongo-driver/v2/bson"

type Address struct {
	Street   string `bson:"street" json:"street"`
	Suburb   string `bson:"suburb" json:"suburb"`
	Postcode string `bson:"postcode" json:"postcode"`
	State    string `bson:"state" json:"state"`
	Country  string `bson:"country" json:"country"`
}

type Booking struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string        `bson:"first_name" json:"first_name"`
	LastName  string        `bson:"last_name" json:"last_name"`
	BookDate  string        `bson:"book_date" json:"book_date"`
	Phone     string        `bson:"phone" json:"phone"`
	Email     string        `bson:"email" json:"email"`
	Address   Address       `bson:"address" json:"address"`
}
