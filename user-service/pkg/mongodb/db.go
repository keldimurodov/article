package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection to MongoDB
func ConnectMongoDB(uri string) (*mongo.Client, error) {
	// Connection settings
	clientOptions := options.Client().ApplyURI(uri)

	// Connection to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Successful connection
	fmt.Println("Successful connection to MongoDB")

	return client, nil
}