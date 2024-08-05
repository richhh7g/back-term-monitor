package mongo_client

import (
	"context"
	"time"

	"github.com/richhh7g/back-term-monitor/pkg/environment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient interface {
	Disconnect(ctx context.Context) error
	FindCollection(name string) *mongo.Collection
}

type MongoClientImpl struct {
	Driver *mongo.Client
	dbName *string
}

func NewMongoClient(ctx context.Context, dbName *string) (MongoClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mongoUrlEnv := environment.Get[string]("MONGO_URL")
	opts := options.Client().
		SetCompressors([]string{"snappy", "zlib", "zstd"}).
		ApplyURI(mongoUrlEnv)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &MongoClientImpl{Driver: client, dbName: dbName}, nil
}

func (c *MongoClientImpl) Disconnect(ctx context.Context) error {
	return c.Driver.Disconnect(ctx)
}

func (c *MongoClientImpl) FindCollection(name string) *mongo.Collection {
	return c.Driver.Database(*c.dbName).Collection(name)
}
