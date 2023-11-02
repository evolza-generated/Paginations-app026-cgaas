package dao

import (
    "context"
	"Paginations/dbConfig"
	"Paginations/dto"

)

func DB_CreateDriver (application *dto.Driver) error {

	_, err := dbConfig.DATABASE.Collection("Drivers").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}