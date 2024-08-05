package brand_usecase

import (
	"context"

	brand_model "github.com/richhh7g/back-term-monitor/internal/domain/model/brand"
	brand_datasource "github.com/richhh7g/back-term-monitor/internal/infra/data/brand"
	"github.com/richhh7g/back-term-monitor/pkg/localization"
)

type CreateBrand interface {
	Exec(ctx context.Context, input *brand_model.CreateBrandInputModel) (*string, error)
}

type CreateBrandImpl struct {
	localization localization.Localization
	datasource   brand_datasource.Brand
}

func NewCreateBrandUseCase(localization localization.Localization, brandDataSource brand_datasource.Brand) CreateBrand {
	return &CreateBrandImpl{
		localization: localization,
		datasource:   brandDataSource,
	}
}

func (u *CreateBrandImpl) Exec(ctx context.Context, input *brand_model.CreateBrandInputModel) (*string, error) {
	_, err := u.datasource.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	message := u.localization.T("brand.success.create", nil)
	return &message, nil
}
