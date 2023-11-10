package dao

import (
	"Paginations/dbConfig"
	"Paginations/dto"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB_FindallCustomer(counter, reqRecords string) (*[]dto.Customer, error) {
	var objects []dto.Customer

	skipCount, err := strconv.Atoi(counter)
	if err != nil {
		return nil, err
	}

	requestedRecordCount, err := strconv.Atoi(reqRecords)
	if err != nil {
		return nil, err
	}

	filter := bson.M{}
	opts := options.Find().SetLimit(int64(requestedRecordCount)).SetSkip(int64(skipCount))

	results, err := dbConfig.DATABASE.Collection("Customers").Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}

	for results.Next(context.Background()) {
		var object dto.Customer
		if err := results.Decode(&object); err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}

	return &objects, nil
}
