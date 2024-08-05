package competitor_datasource

import (
	"context"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
)

type Competitor interface {
	Create(ctx context.Context, input *competitor_model.CreateCompetitorInputModel) ([]*competitor_model.CompetitorBaseModel, error)
}

type CompetitorImpl struct {
	competitorDBDataSource   CompetitorDB
	competitorHttpDataSource CompetitorHttp
}

func NewCompetitorDataSource(competitorHttpDataSource CompetitorHttp, competitorDBDataSource CompetitorDB) Competitor {
	return &CompetitorImpl{
		competitorDBDataSource:   competitorDBDataSource,
		competitorHttpDataSource: competitorHttpDataSource,
	}
}

func (d *CompetitorImpl) Create(ctx context.Context, input *competitor_model.CreateCompetitorInputModel) ([]*competitor_model.CompetitorBaseModel, error) {
	searchTermsModel, err := d.competitorHttpDataSource.FindTerm(ctx, &competitor_model.FindTermCompetitorHttpInputModel{
		Term:   input.Term,
		City:   input.City,
		Device: input.Device,
	})
	if err != nil {
		return nil, err
	}

	var competitorsModel []*competitor_model.CompetitorBaseModel

	for _, searchTermModel := range searchTermsModel {
		competitorDbModel, err := d.competitorDBDataSource.Create(ctx, &competitor_model.CreateCompetitorDBInputModel{
			Link:        searchTermModel.Link,
			City:        searchTermModel.City,
			Term:        searchTermModel.Term,
			Device:      searchTermModel.Device,
			BrandID:     input.BrandID,
			ProcessedAt: searchTermModel.ProcessedAt,
		})
		if err != nil {
			return nil, err
		}

		competitorsModel = append(competitorsModel, competitorDbModel)
	}

	return competitorsModel, nil
}
