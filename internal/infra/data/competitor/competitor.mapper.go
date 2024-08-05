package competitor_datasource

import (
	"strings"

	competitor_model "github.com/richhh7g/back-term-monitor/internal/domain/model/competitor"
	serpapi_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/serpapi"
)

var allLocations = []serpapi_client.Location{
	serpapi_client.SaoPaulo,
	serpapi_client.RioDeJaneiro,
	serpapi_client.Brasilia,
	serpapi_client.Salvador,
	serpapi_client.Fortaleza,
	serpapi_client.BeloHorizonte,
	serpapi_client.Manaus,
	serpapi_client.Curitiba,
	serpapi_client.Recife,
	serpapi_client.PortoAlegre,
}

var allDevicesClient = []serpapi_client.Device{
	serpapi_client.Mobile,
	serpapi_client.Tablet,
	serpapi_client.Desktop,
}

func mapCityToLocation(city competitor_model.City) *serpapi_client.Location {
	for _, location := range allLocations {
		if strings.Contains(string(location), string(city)) {
			return &location
		}
	}

	return nil
}

func mapDeviceToDeviceClient(device competitor_model.Device) *serpapi_client.Device {
	for _, deviceClient := range allDevicesClient {
		lowerDevice := strings.ToLower(string(device))
		lowerDeviceClient := strings.ToLower(string(deviceClient))

		if strings.Contains(lowerDeviceClient, lowerDevice) {
			return &deviceClient
		}
	}

	return nil
}
