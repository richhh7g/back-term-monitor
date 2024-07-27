package mongo_repository

import (
	"context"
	"time"

	mongo_client "github.com/richhh7g/term-alarms/infra/data/client/mongo"
	mongo_document "github.com/richhh7g/term-alarms/infra/data/client/mongo/document"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "alarm"

type AlarmRepository interface {
	Create(ctx context.Context, document *mongo_document.Alarm, opts ...*options.InsertOneOptions) (*mongo_document.Alarm, error)
	FindOneByEmail(ctx context.Context, email string, opts ...*options.FindOneOptions) (*mongo_document.Alarm, error)
}

type AlarmRepositoryImpl struct {
	client          mongo_client.MongoClient
	alarmCollection *mongo.Collection
}

func NewAlarmRepository(client mongo_client.MongoClient) AlarmRepository {
	alarmCollection := client.FindCollection(collectionName)

	return &AlarmRepositoryImpl{client: client, alarmCollection: alarmCollection}
}

func (r *AlarmRepositoryImpl) Create(ctx context.Context, document *mongo_document.Alarm, opts ...*options.InsertOneOptions) (*mongo_document.Alarm, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.alarmCollection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (r *AlarmRepositoryImpl) FindOneByEmail(ctx context.Context, email string, opts ...*options.FindOneOptions) (*mongo_document.Alarm, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var document mongo_document.Alarm
	err := r.alarmCollection.FindOne(ctx, bson.M{"email": email}, opts...).Decode(&document)

	if err == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}

	if err != nil {
		return nil, err
	}

	return &document, nil
}
