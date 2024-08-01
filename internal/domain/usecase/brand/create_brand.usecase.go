package brand_usecase

import (
	"context"

	"github.com/richhh7g/term-alarms/internal/domain/model"
	brand_datasource "github.com/richhh7g/term-alarms/internal/infra/data/brand"
	"github.com/richhh7g/term-alarms/pkg/localization"
)

type CreateBrand interface {
	Exec(ctx context.Context, input *model.CreateBrandInputModel) (*string, error)
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

func (u *CreateBrandImpl) Exec(ctx context.Context, input *model.CreateBrandInputModel) (*string, error) {
	_, err := u.datasource.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	message := u.localization.T("brand.success.create", nil)
	return &message, nil
}
