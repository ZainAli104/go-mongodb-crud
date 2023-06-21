package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://zainfiverr44:ra4Trr0tcn6RyLt9@cluster0.sixhmey.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("test")
	log.Println("Connected to MongoDB!")
}
