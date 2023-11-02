package dbConfig

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ConnectToMongoDB() {
	fmt.Println("Connecting to mongo cluster")

	// Create a context
	ctx := context.Background()

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}

	//List Available Databases in the Cluster
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to mongo cluster")
	}

	DATABASE = client.Database(DATABASE_NAME)

	fmt.Println(databases)

}
