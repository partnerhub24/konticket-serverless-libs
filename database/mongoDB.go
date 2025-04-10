package database

import (
	"context"
	"fmt"
	"github.com/partnerhub24/konticket-serverless-libs/environment"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

func InitMongoDB(mongoEnv environment.MongoDBEnvironment) (*mongo.Client, error) {
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=konticket-queue-dev",
		mongoEnv.MongoUser,
		mongoEnv.MongoPass,
		mongoEnv.MongoHost,
		mongoEnv.MongoName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri).SetMaxPoolSize(5). // limit concurrent connections
									SetMinPoolSize(1).
									SetConnectTimeout(10 * time.Second).SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(clientOpts)

	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	return client, nil
}
