package competitor_usecase

import (
	"context"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
	brand_datasource "github.com/richhh7g/back-term-monitor/internal/infra/data/brand"
	competitor_datasource "github.com/richhh7g/back-term-monitor/internal/infra/data/competitor"
	"golang.org/x/exp/rand"
)

type ProcessCompetitors interface {
	Exec(ctx context.Context) ([]*competitor_model.CompetitorTermBaseModel, error)
}

type ProcessCompetitorsImpl struct {
	brandDataSource      brand_datasource.Brand
	competitorDataSource competitor_datasource.Competitor
}

func NewProcessCompetitorsUseCase(brandDataSource brand_datasource.Brand, competitorDataSource competitor_datasource.Competitor) ProcessCompetitors {
	return &ProcessCompetitorsImpl{
		brandDataSource:      brandDataSource,
		competitorDataSource: competitorDataSource,
	}
}

func (u *ProcessCompetitorsImpl) Exec(ctx context.Context) ([]*competitor_model.CompetitorTermBaseModel, error) {
	brandDb, err := u.brandDataSource.FindProcessingBrand(ctx)
	if err != nil {
		return nil, err
	}
	if brandDb == nil {
		return nil, nil
	}

	var deviceSelected = u.randomDevice()

	var competitorsTermsBaseModel []*competitor_model.CompetitorTermBaseModel
	for _, term := range brandDb.Terms {
		competitorsTermsBaseModel = append(competitorsTermsBaseModel, u.loopForCities(ctx, term, brandDb.ID, deviceSelected))
	}

	isCompleted, err := u.brandDataSource.UpdateSuccess(ctx, brandDb.ID)
	if err != nil {
		return nil, err
	}

	if !isCompleted {
		return nil, nil
	}

	return competitorsTermsBaseModel, nil
}

func (u *ProcessCompetitorsImpl) randomDevice() competitor_model.Device {
	var devices []competitor_model.Device = []competitor_model.Device{
		competitor_model.Desktop,
		competitor_model.Mobile,
		competitor_model.Tablet,
	}

	return devices[rand.Intn(len(devices))]
}

func (u *ProcessCompetitorsImpl) loopForCities(ctx context.Context, term string, brandId string, device competitor_model.Device) *competitor_model.CompetitorTermBaseModel {
	var cities []competitor_model.City = []competitor_model.City{
		competitor_model.SaoPaulo,
		competitor_model.RioDeJaneiro,
		competitor_model.Brasilia,
		competitor_model.Salvador,
		competitor_model.Fortaleza,
		competitor_model.BeloHorizonte,
		competitor_model.Manaus,
		competitor_model.Curitiba,
		competitor_model.Recife,
		competitor_model.PortoAlegre,
	}

	var competitorsCities []*competitor_model.CompetitorCityWithLinksBaseModel
	competitorTermBaseModel := &competitor_model.CompetitorTermBaseModel{
		Term: term,
	}

	for _, city := range cities {
		competitorBaseModel, err := u.competitorDataSource.Create(ctx, &competitor_model.CreateCompetitorInputModel{
			Term:    term,
			City:    city,
			Device:  device,
			BrandID: brandId,
		})
		if err != nil {
			continue
		}

		var links []string
		for _, competitor := range competitorBaseModel {
			links = append(links, competitor.Link)
		}

		competitorCityWithLinksBaseModel := &competitor_model.CompetitorCityWithLinksBaseModel{
			Name:  string(city),
			Links: links,
		}

		competitorsCities = append(competitorsCities, competitorCityWithLinksBaseModel)
	}

	competitorTermBaseModel.Cities = competitorsCities

	return competitorTermBaseModel
}
