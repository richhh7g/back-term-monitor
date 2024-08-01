package brand_datasource

import (
	"context"

	"github.com/richhh7g/term-alarms/internal/domain/model"
	mongo_document "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/document"
	mongo_repository "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/repository"
)

type Brand interface {
	Create(ctx context.Context, input *model.CreateBrandInputModel) (bool, error)
}

type BrandImpl struct {
	repository mongo_repository.BrandRepository
}

func NewBrandDataSource(brandRepository mongo_repository.BrandRepository) Brand {
	return &BrandImpl{
		repository: brandRepository,
	}
}

func (d *BrandImpl) Create(ctx context.Context, input *model.CreateBrandInputModel) (bool, error) {
	_, err := d.repository.Create(ctx, &mongo_repository.CreateBrandParams{
		Email:      input.Email,
		Status:     mongo_document.BrandPendingStatus,
		BrandTerms: input.Terms,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
