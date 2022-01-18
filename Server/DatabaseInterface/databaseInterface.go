package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const localURI = "mongodb://localhost:27017"

func mongodbServiceConnect(uri string) (mongo.Client, error) {

	// Makes a 'client' connection to the MongoDB server service, passed as the string URI. Logs fatal error if process somehow incomplete.
	mongodbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Conducts an initial ping operation to check if connection to MongoDB server service successful.
	err = mongodbClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	// Assuming above steps are succesful, notify server terminal.
	fmt.Println("Successfully connected to local MongoDB server service.")

	return *mongodbClient, nil
}

func mongodbServiceDisconnect(client mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Print(err)
	}
}

func main() {
	// Connect to MongoDB service
	client, err := mongodbServiceConnect(localURI)
	if err != nil {
		log.Print(err)
	}

	// Create new members in database
	// NewMember(client, "Brandon Fisher", "BigApple123!", Enums.LowTier)
	// NewMember(client, "Sarah Reynolds", "PearlyGates", Enums.AdminTier)

	// Read from members
	var members []Member
	members, err = GetAllMembers(client)
	if err != nil {
		log.Print(err)
	}

	// Read from passed members
	fmt.Println(members)

	// Disconnect from MongoDB service
	mongodbServiceDisconnect(client)
}
