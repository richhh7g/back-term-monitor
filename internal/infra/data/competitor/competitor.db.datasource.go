package competitor_datasource

import (
	"context"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
	mongo_repository "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompetitorDB interface {
	Create(ctx context.Context, input *competitor_model.CreateCompetitorDBInputModel) (*competitor_model.CompetitorBaseModel, error)
}

type CompetitorDBImpl struct {
	repository mongo_repository.CompetitorRepository
}

func NewCompetitorDBDataSource(competitorRepository mongo_repository.CompetitorRepository) CompetitorDB {
	return &CompetitorDBImpl{
		repository: competitorRepository,
	}
}

func (d *CompetitorDBImpl) Create(ctx context.Context, input *competitor_model.CreateCompetitorDBInputModel) (*competitor_model.CompetitorBaseModel, error) {
	brandID, err := primitive.ObjectIDFromHex(input.BrandID)
	if err != nil {
		return nil, err
	}

	processedAt := primitive.Timestamp{
		T: uint32(input.ProcessedAt.Unix()),
		I: 0,
	}

	competitorDB, err := d.repository.Create(ctx, &mongo_repository.CreateCompetitorParams{
		Term:     input.Term,
		Device:   string(input.Device),
		Domain:   input.Link,
		BrandID:  brandID,
		FoundAt:  processedAt,
		Location: string(input.City),
	})
	if err != nil {
		return nil, err
	}

	if competitorDB == nil {
		return nil, nil
	}

	return &competitor_model.CompetitorBaseModel{
		ID:      competitorDB.ID.Hex(),
		Link:    input.Link,
		City:    input.City,
		Term:    input.Term,
		Device:  input.Device,
		BrandID: input.BrandID,
	}, nil
}
