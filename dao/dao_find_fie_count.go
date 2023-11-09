package dao

import (
	"Paginations/dbConfig"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func FindFileCountDao() (int64, error) {
	collection := dbConfig.DATABASE.Collection("Customers")
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return 0, fmt.Errorf("Error while counting documents: %v", err)
	}
	return count, nil
}
