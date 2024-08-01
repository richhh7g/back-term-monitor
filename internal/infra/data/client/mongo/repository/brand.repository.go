package mongo_repository

import (
	"context"
	"time"

	mongo_client "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo"
	mongo_document "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/document"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateBrandParams struct {
	Email      string
	Status     string
	BrandTerms []string
}

const brandCollectionName = "brand"

type BrandRepository interface {
	Create(ctx context.Context, input *CreateBrandParams, opts ...*options.InsertOneOptions) (*mongo_document.Brand, error)
}

type BrandRepositoryImpl struct {
	client          mongo_client.MongoClient
	brandCollection *mongo.Collection
}

func NewBrandRepository(client mongo_client.MongoClient) BrandRepository {
	brandCollection := client.FindCollection(brandCollectionName)

	return &BrandRepositoryImpl{client: client, brandCollection: brandCollection}
}

func (r *BrandRepositoryImpl) Create(ctx context.Context, input *CreateBrandParams, opts ...*options.InsertOneOptions) (*mongo_document.Brand, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	document := mongo_document.Brand{
		Email:      input.Email,
		Status:     input.Status,
		BrandTerms: input.BrandTerms,
	}

	_, err := r.brandCollection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
