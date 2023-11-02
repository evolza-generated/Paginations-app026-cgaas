package dao

import (
	"Paginations/dbConfig"
	"Paginations/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallDriver () (*[]dto.Driver, error) {
	var objects []dto.Driver
	results, err := dbConfig.DATABASE.Collection("Drivers").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Driver")
	}
	for results.Next(context.Background()) {
		var object dto.Driver
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Driver")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
