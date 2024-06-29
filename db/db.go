package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client *mongo.Client
}

func NewStorage(uri string) (*Storage, error) {
	client, err := createNewStorage(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Database %w", err)
	}

	// check connection
	err = checkConnection(client)
	if err != nil {
		return nil, fmt.Errorf("error checking connection with database %w", err)
	}

	return &Storage{
		Client: client,
	}, nil
}

func checkConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}

func createNewStorage(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	return client, err
}
