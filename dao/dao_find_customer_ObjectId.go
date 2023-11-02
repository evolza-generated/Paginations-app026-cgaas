package dao

import (
	"Paginations/dbConfig"
	"Paginations/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindCustomerbyObjectId (objectid string) (*dto.Customer, error) {
	var object dto.Customer

	err := dbConfig.DATABASE.Collection("Customers").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Customer")
    }
	if object == (dto.Customer{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
