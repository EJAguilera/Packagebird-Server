package main

import (
	"Server/Enums"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Member struct {
	Name     string               `bson:"name"`
	Password string               `bson:"password"`
	Employed bool                 `bson:"is_employed"`
	Tier     Enums.PrivilegeLevel `bson:"privilege_level"`
}

func NewMember(client mongo.Client, name string, password string, tier Enums.PrivilegeLevel) {
	// First establish connection to database and collection
	collection := client.Database("packagebird").Collection("members")

	// Check if Tier is valid
	if !Enums.IsTier(tier) {
		log.Print("Tier assigned to new member is not valid. Please try again with valid tier.")
		return
	}

	// If valid, proceed to new member creation
	newMember := Member{Name: name, Password: password, Tier: tier, Employed: true}
	_, err := collection.InsertOne(context.TODO(), newMember)
	if err != nil {
		log.Print(err)
	}
}

func GetAllMembers(client mongo.Client) ([]Member, error) {
	// First establish connection to database and collection
	collection := client.Database("packagebird").Collection("members")

	// Retrieve all entries in collection
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var memberResults []Member
	err = cursor.All(context.TODO(), &memberResults)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	// Print for veracity
	fmt.Println(memberResults)

	// Pass results back to caller
	return memberResults, nil
}
