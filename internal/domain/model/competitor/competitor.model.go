package competitor_model

import "time"

type City string

const (
	SaoPaulo      City = "Sao Paulo"
	RioDeJaneiro  City = "Rio de Janeiro"
	Brasilia      City = "Brasilia"
	Salvador      City = "Salvador"
	Fortaleza     City = "Fortaleza"
	BeloHorizonte City = "Belo Horizonte"
	Manaus        City = "Manaus"
	Curitiba      City = "Curitiba"
	Recife        City = "Recife"
	PortoAlegre   City = "Porto Alegre"
)

type Device string

const (
	Tablet  Device = "Tablet"
	Mobile  Device = "Mobile"
	Desktop Device = "Desktop"
)

type CompetitorBaseModel struct {
	ID      string
	Link    string
	City    City
	Term    string
	Device  Device
	BrandID string
}

type CreateCompetitorInputModel struct {
	City    City
	Term    string
	Device  Device
	BrandID string
}

type CreateCompetitorDBInputModel struct {
	Link        string
	City        City
	Term        string
	Device      Device
	BrandID     string
	ProcessedAt time.Time
}

type FindTermCompetitorHttpInputModel struct {
	Term   string
	City   City
	Device Device
}

type SearchTermDataModel struct {
	City             City
	Link             string
	Term             string
	Device           Device
	ProcessedAt      time.Time
	BrowserSearchUrl string
}
