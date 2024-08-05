package mongo_repository

import (
	"context"
	"time"

	mongo_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo"
	mongo_document "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/document"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateBrandParams struct {
	Email      string
	Status     string
	BrandTerms []string
}

type FindOneBrandByParams struct {
	ID     string
	Email  string
	Status string
}

type UpdateBrandParams struct {
	Email      string               `bson:"email"`
	Status     string               `bson:"status"`
	Results    []primitive.ObjectID `bson:"results"`
	BrandTerms []string             `bson:"brand_terms"`
}

const brandCollectionName = "brand"

type BrandRepository interface {
	Create(ctx context.Context, input *CreateBrandParams, opts ...*options.InsertOneOptions) (*mongo_document.Brand, error)
	FindOneBy(ctx context.Context, input *FindOneBrandByParams, opts ...*options.FindOneOptions) (*mongo_document.Brand, error)
	Update(ctx context.Context, brandId primitive.ObjectID, input *UpdateBrandParams, opts ...*options.UpdateOptions) (bool, error)
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

func (r *BrandRepositoryImpl) FindOneBy(ctx context.Context, input *FindOneBrandByParams, opts ...*options.FindOneOptions) (*mongo_document.Brand, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var document mongo_document.Brand

	filter := bson.M{}
	if input.ID != "" {
		filter["_id"] = input.ID
	}
	if input.Email != "" {
		filter["email"] = input.Email
	}
	if input.Status != "" {
		filter["status"] = input.Status
	}

	err := r.brandCollection.FindOne(ctx, filter, opts...).Decode(&document)

	if err == mongo.ErrNoDocuments {
		return nil, mongo.ErrNoDocuments
	}

	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *BrandRepositoryImpl) Update(ctx context.Context, brandId primitive.ObjectID, input *UpdateBrandParams, opts ...*options.UpdateOptions) (bool, error) {
	filter := bson.M{"_id": brandId}

	updateFields := bson.M{"$set": input}
	result, err := r.brandCollection.UpdateOne(ctx, filter, updateFields, opts...)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, mongo.ErrNoDocuments
	}

	return true, nil
}
