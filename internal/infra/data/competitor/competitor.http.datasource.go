package competitor_datasource

import (
	"context"
	"time"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
	serpapi_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/serpapi"
)

type CompetitorHttp interface {
	FindTerm(ctx context.Context, input *competitor_model.FindTermCompetitorHttpInputModel) ([]*competitor_model.SearchTermDataModel, error)
}

type CompetitorHttpImpl struct {
	client serpapi_client.SerpApi
}

func NewCompetitorHttpDataSource(client serpapi_client.SerpApi) CompetitorHttp {
	return &CompetitorHttpImpl{
		client: client,
	}
}

func (d *CompetitorHttpImpl) FindTerm(ctx context.Context, input *competitor_model.FindTermCompetitorHttpInputModel) ([]*competitor_model.SearchTermDataModel, error) {
	mappedLocation := mapCityToLocation(input.City)
	mappedDevice := mapDeviceToDeviceClient(input.Device)

	response, err := d.client.GoogleSearch(&serpapi_client.GoogleSearchRequest{
		Device:   *mappedDevice,
		Location: *mappedLocation,
		Text:     input.Term,
	})
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	var searchTerms []*competitor_model.SearchTermDataModel

	for _, advertisement := range response.Ads {
		const dateLayout = "2006-01-02 15:04:05 MST"
		processedAt, err := time.Parse(dateLayout, response.SearchMetadata.ProcessedAt)
		if err != nil {
			return nil, err
		}

		link := advertisement.Link
		if link == "" {
			link = mapSourceToLink(advertisement.Source)
		}

		searchTerms = append(searchTerms, &competitor_model.SearchTermDataModel{
			City:             input.City,
			Link:             link,
			Term:             input.Term,
			Device:           input.Device,
			ProcessedAt:      processedAt,
			BrowserSearchUrl: response.SearchMetadata.GoogleUrl,
		})
	}

	return searchTerms, nil
}
