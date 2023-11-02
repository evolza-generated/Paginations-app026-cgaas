package dao

import (
	"Paginations/dbConfig"
	"Paginations/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallCustomer () (*[]dto.Customer, error) {
	var objects []dto.Customer
	results, err := dbConfig.DATABASE.Collection("Customers").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Customer")
	}
	for results.Next(context.Background()) {
		var object dto.Customer
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Customer")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
