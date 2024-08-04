package serpapi_client

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/richhh7g/back-term-monitor/pkg/environment"
)

type SerpApi interface {
	GoogleAds(params *GoogleAdsSearchRequest) (*GoogleAdsResponse, error)
	GoogleSearch(params *GoogleSearchRequest) (*GoogleSearchResponse, error)
}

type SerpApiImpl struct {
	driver *resty.Client
	apiKey string
}

func NewSerpApiClient() SerpApi {
	client := resty.New()
	client.BaseURL = environment.Get[string]("SERPAPI_BASE_URL")
	client.Header.Add("Content-Type", "application/json")

	apiKey := environment.Get[string]("SERPAPI_API_KEY")

	return &SerpApiImpl{
		driver: client,
		apiKey: apiKey,
	}
}

func (c *SerpApiImpl) GoogleAds(params *GoogleAdsSearchRequest) (*GoogleAdsResponse, error) {
	var response GoogleAdsResponse

	result, err := c.driver.R().
		SetQueryParams(map[string]string{
			"q":        params.Text,
			"engine":   "google_ads_transparency_center",
			"location": strconv.Itoa(int(params.Region)),
			"api_key":  c.apiKey,
		}).
		SetResult(&response).
		Get("/search.json")

	if result.StatusCode() != 200 {
		return nil, fmt.Errorf("response status code: %d", result.StatusCode())
	}

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *SerpApiImpl) GoogleSearch(params *GoogleSearchRequest) (*GoogleSearchResponse, error) {
	var response GoogleSearchResponse

	result, err := c.driver.R().
		SetQueryParams(map[string]string{
			"q":             params.Text,
			"location":      string(params.Location),
			"device":        string(params.Device),
			"engine":        "google",
			"google_domain": "google.com.br",
			"gl":            "br",
			"hl":            "pt",
			"safe":          "off",
			"api_key":       c.apiKey,
		}).
		SetResult(&response).
		Get("/search.json")

	if result.StatusCode() != 200 {
		return nil, fmt.Errorf("response status code: %d", result.StatusCode())
	}

	if err != nil {
		return nil, err
	}

	if len(response.Ads) == 0 {
		return nil, nil
	}

	return &response, nil
}
