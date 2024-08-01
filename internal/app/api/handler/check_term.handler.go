package handler

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/eduardolat/goeasyi18n"
	"github.com/go-playground/validator/v10"
	"github.com/richhh7g/term-alarms/internal/app/api/response"
	"github.com/richhh7g/term-alarms/internal/domain/model"
	brand_usecase "github.com/richhh7g/term-alarms/internal/domain/usecase/brand"
	brand_datasource "github.com/richhh7g/term-alarms/internal/infra/data/brand"
	mongo_client "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo"
	mongo_repository "github.com/richhh7g/term-alarms/internal/infra/data/client/mongo/repository"
	"github.com/richhh7g/term-alarms/pkg/environment"
	"github.com/richhh7g/term-alarms/pkg/localization"
)

func NewCheckTermHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	var bodyParsed CheckTermBody
	if err := json.NewDecoder(r.Body).Decode(&bodyParsed); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	// TODO: Criar middleware de validação
	err := validator.New().Struct(bodyParsed)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	// TODO: Adicionar injeção de dependência
	databaseNameEnv := environment.Get[string]("MONGO_DB")
	mongoClient, _ := mongo_client.NewMongoClient(ctx, &databaseNameEnv)
	defer mongoClient.Disconnect(ctx)

	brandRepository := mongo_repository.NewBrandRepository(mongoClient)
	brandDatasource := brand_datasource.NewBrandDataSource(brandRepository)
	localizationService := localization.NewLocalization(goeasyi18n.NewI18n())
	localizationService.AddLanguages(map[localization.Language]string{
		localization.EN_US: path.Join("pkg", "localization", "locale", "en_us.locale.yml"),
		localization.PT_BR: path.Join("pkg", "localization", "locale", "pt_br.locale.yml"),
	})

	brandUseCase := brand_usecase.NewCreateBrandUseCase(localizationService, brandDatasource)
	result, err := brandUseCase.Exec(ctx, &model.CreateBrandInputModel{
		Email: bodyParsed.Email,
		Terms: bodyParsed.Termos,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		message := localizationService.T("server.error.internal", nil)
		json.NewEncoder(w).Encode(map[string]string{"error": message})

		return
	}

	response.NewSuccess(map[string]string{"text": *result}, http.StatusOK).Send(w)
}
