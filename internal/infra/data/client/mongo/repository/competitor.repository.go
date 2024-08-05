package mongo_repository

import (
	"context"
	"time"

	mongo_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo"
	mongo_document "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/document"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateCompetitorParams struct {
	Term     string
	Device   string
	Domain   string
	FoundAt  primitive.Timestamp
	BrandID  primitive.ObjectID
	Location string
}

const competitorCollectionName = "competitor"

type CompetitorRepository interface {
	Create(ctx context.Context, input *CreateCompetitorParams, opts ...*options.InsertOneOptions) (*mongo_document.Competitor, error)
}

type CompetitorRepositoryImpl struct {
	client               mongo_client.MongoClient
	competitorCollection *mongo.Collection
}

func NewCompetitorRepository(client mongo_client.MongoClient) CompetitorRepository {
	competitorCollection := client.FindCollection(competitorCollectionName)

	return &CompetitorRepositoryImpl{client: client, competitorCollection: competitorCollection}
}

func (r *CompetitorRepositoryImpl) Create(ctx context.Context, input *CreateCompetitorParams, opts ...*options.InsertOneOptions) (*mongo_document.Competitor, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	document := mongo_document.Competitor{
		Term:     input.Term,
		Device:   input.Device,
		Domain:   input.Domain,
		FoundAt:  input.FoundAt,
		BrandID:  input.BrandID,
		Location: input.Location,
	}

	_, err := r.competitorCollection.InsertOne(ctx, document, opts...)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
