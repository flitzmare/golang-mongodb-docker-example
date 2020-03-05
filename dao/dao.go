package dao

import (
	"context"
	"echo-mongo/model"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDBConnection() (*mongo.Collection, error) {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mydb").Collection("persons")

	return collection, nil
}

func InsertUser(u *model.User) {
	collection, err := GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := collection.InsertOne(context.TODO(), u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document of user: ", insertResult.InsertedID)
}
