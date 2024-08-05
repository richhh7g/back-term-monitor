package schedule_job

import (
	"context"

	competitor_usecase "github.com/richhh7g/back-term-monitor/internal/domain/usecase/competitor"
	brand_datasource "github.com/richhh7g/back-term-monitor/internal/infra/data/brand"
	mongo_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo"
	mongo_repository "github.com/richhh7g/back-term-monitor/internal/infra/data/client/mongo/repository"
	serpapi_client "github.com/richhh7g/back-term-monitor/internal/infra/data/client/serpapi"
	competitor_datasource "github.com/richhh7g/back-term-monitor/internal/infra/data/competitor"
	"github.com/richhh7g/back-term-monitor/pkg/environment"
)

type processCompetitorsImpl struct{}

func NewProcessCompetitors() Job {
	return &processCompetitorsImpl{}
}

func (j *processCompetitorsImpl) Run() {
	ctx := context.Background()

	// TODO: Adicionar injeção de dependência
	databaseNameEnv := environment.Get[string]("MONGO_DB")
	mongoClient, _ := mongo_client.NewMongoClient(ctx, &databaseNameEnv)
	defer mongoClient.Disconnect(ctx)

	brandDataSource := brand_datasource.NewBrandDbDataSource(mongo_repository.NewBrandRepository(mongoClient))

	competitorRepository := mongo_repository.NewCompetitorRepository(mongoClient)
	competitorDBDataSource := competitor_datasource.NewCompetitorDBDataSource(competitorRepository)

	client := serpapi_client.NewSerpApiClient()
	competitorHttpDataSource := competitor_datasource.NewCompetitorHttpDataSource(client)
	competitorDatasource := competitor_datasource.NewCompetitorDataSource(competitorHttpDataSource, competitorDBDataSource)

	processCompetitorsUseCase := competitor_usecase.NewProcessCompetitorsUseCase(brandDataSource, competitorDatasource)
	processCompetitorsUseCase.Exec(ctx)
}
