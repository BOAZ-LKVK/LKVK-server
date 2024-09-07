package main

import (
	recommendation_api "github.com/BOAZ-LKVK/LKVK-server/api/recommendation"
	sample_api "github.com/BOAZ-LKVK/LKVK-server/api/sample"
	"github.com/BOAZ-LKVK/LKVK-server/pkg/fx/fiberfx"
	"github.com/BOAZ-LKVK/LKVK-server/pkg/fx/gormfx"
	"github.com/BOAZ-LKVK/LKVK-server/pkg/fx/zapfx"
	recommendation_repository "github.com/BOAZ-LKVK/LKVK-server/repository/recommendation"
	sample_repository "github.com/BOAZ-LKVK/LKVK-server/repository/sample"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: log} }),
		fx.Provide(
			zapfx.NewZapLogger,
		),
		fx.Provide(
			sample_repository.NewSampleRepository,
			fiberfx.AsAPIController(sample_api.NewSampleAPIHandler),
			fiberfx.AsAPIController(recommendation_api.NewRecommendationAPIController),
			recommendation_repository.NewRestaurantRecommendationRepository,
			recommendation_repository.NewRestaurantRecommendationRequestRepository,
		),
		fiberfx.Module,
		gormfx.Module,
	).Run()
}
