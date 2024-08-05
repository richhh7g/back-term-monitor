package brand_model

type Status string

const (
	Pending Status = "Pending"
	Success Status = "Success"
)

type BrandBase struct {
	ID      string
	Email   string
	Status  string
	Results []string
	Terms   []string
}

type CreateBrandInputModel struct {
	Email  string
	Terms  []string
	Status Status
}
