package brand_datasource

import (
	"context"

	brand_model "github.com/richhh7g/back-term-monitor/internal/domain/model/brand"
	mongo_document "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/document"
	mongo_repository "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand interface {
	Create(ctx context.Context, input *brand_model.CreateBrandInputModel) (bool, error)
	FindProcessingBrand(ctx context.Context) (*brand_model.BrandBase, error)
	UpdateSuccess(ctx context.Context, id string) (bool, error)
}

type BrandImpl struct {
	repository mongo_repository.BrandRepository
}

func NewBrandDbDataSource(brandRepository mongo_repository.BrandRepository) Brand {
	return &BrandImpl{
		repository: brandRepository,
	}
}

func (d *BrandImpl) Create(ctx context.Context, input *brand_model.CreateBrandInputModel) (bool, error) {
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

func (d *BrandImpl) FindProcessingBrand(ctx context.Context) (*brand_model.BrandBase, error) {
	brandDb, err := d.repository.FindOneBy(ctx, &mongo_repository.FindOneBrandByParams{
		Status: mongo_document.BrandPendingStatus,
	})

	if err != nil {
		return nil, err
	}

	if brandDb == nil {
		return nil, nil
	}

	var results []string
	for _, id := range brandDb.Results {
		results = append(results, id.Hex())
	}

	return &brand_model.BrandBase{
		ID:      brandDb.ID.Hex(),
		Email:   brandDb.Email,
		Status:  brandDb.Status,
		Terms:   brandDb.BrandTerms,
		Results: results,
	}, nil
}

func (d *BrandImpl) UpdateSuccess(ctx context.Context, id string) (bool, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	return d.repository.Update(ctx, objectId, &mongo_repository.UpdateBrandParams{
		Status: mongo_document.BrandSuccessStatus,
	})
}
