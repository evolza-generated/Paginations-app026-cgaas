package dao

import (
	"Paginations/dbConfig"
	"Paginations/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindDriverbyObjectId (objectid string) (*dto.Driver, error) {
	var object dto.Driver

	err := dbConfig.DATABASE.Collection("Drivers").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Driver")
    }
	if object == (dto.Driver{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
