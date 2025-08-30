package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lucasschilin/crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	DATABASE_URI  = "DATABASE_URI"
	DATABASE_NAME = "DATABASE_NAME"
)

func NewConnection(ctx context.Context) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv(DATABASE_URI)))
	logger.Info("Database connection closed successfully")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	logger.Info("Database connection closed successfully")

	return client.Database(os.Getenv(DATABASE_NAME)), nil
}
