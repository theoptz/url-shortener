package mongodb

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database
var err error

// GetClient can possibly return nil
// Please be sure that you have created instance on application startup via "Start" function
func GetClient() *mongo.Client {
	return client
}

// GetDatabase can possibly return nil
// Please be sure that you have created instance on application startup via "Start" function
func GetDatabase() *mongo.Database {
	return database
}

func Close(ctx context.Context) {
	_ = client.Disconnect(ctx)
}

func Connect(ctx context.Context, uri string, databaseName string) (*mongo.Client, error) {
	var once sync.Once
	once.Do(func() {
		client, err = mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			return
		}

		err = client.Connect(ctx)
		if err != nil {
			return
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return
		}

		database = client.Database(databaseName)
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}
