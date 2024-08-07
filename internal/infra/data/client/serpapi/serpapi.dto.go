package serpapi_client

type Country int

// Source: https://serpapi.com/google-ads-transparency-center-regions
const (
	CountryBrazil Country = 2076
)

type Location string

const (
	SaoPaulo      Location = "Sao Paulo, State of Sao Paulo, Brazil"
	RioDeJaneiro  Location = "Rio de Janeiro, State of Rio de Janeiro, Brazil"
	Brasilia      Location = "Brasilia, Federal District, Brazil"
	Salvador      Location = "Salvador, Salvador, State of Bahia, Brazil"
	Fortaleza     Location = "Fortaleza, Ceara, Brazil"
	BeloHorizonte Location = "Belo Horizonte, State of Minas Gerais,Brazil"
	Manaus        Location = "Manaus, State of Amazonas, Brazil"
	Curitiba      Location = "Curitiba, State of Parana, Brazil"
	Recife        Location = "Recife, State of Pernambuco, Brazil"
	PortoAlegre   Location = "Porto Alegre, State of Rio Grande do Sul, Brazil"
)

type Device string

const (
	Desktop Device = "desktop"
	Tablet  Device = "tablet"
	Mobile  Device = "mobile"
)

type searchMetadata struct {
	Status                         string `json:"status"`
	ProcessedAt                    string `json:"processed_at"`
	GoogleUrl                      string `json:"google_url"`
	GoogleAdsTransparencyCenterUrl string `json:"google_ads_transparency_center_url"`
}

type searchParameters struct {
	Text string `json:"text"`
}

type adCreative struct {
	Advertiser   string `json:"advertiser"`
	TargetDomain string `json:"target_domain"`
	DetailsLink  string `json:"details_link"`
}

type GoogleAdsSearchRequest struct {
	Region Country // Region
	Text   string  // Query to search
}

type GoogleAdsResponse struct {
	SearchMetadata searchMetadata   `json:"search_metadata"`
	SearchParams   searchParameters `json:"search_parameters"`
	AdCreatives    []adCreative     `json:"ad_creatives"`
}

type advertisement struct {
	Link         string `json:"link"`
	Title        string `json:"title"`
	Source       string `json:"source"`
	Description  string `json:"description"`
	TrackingLink string `json:"tracking_link"`
}

type GoogleSearchRequest struct {
	Text     string   // Query to search
	Location Location // Location
	Device   Device   // Device
}

type GoogleSearchResponse struct {
	SearchMetadata searchMetadata  `json:"search_metadata"`
	Ads            []advertisement `json:"ads"`
}
