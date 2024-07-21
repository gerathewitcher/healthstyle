package app

import (
	"context"
	"log"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/config"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/repository"
	rep "github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/service"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/closer"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db/pg"

	nutrition "github.com/gerathewitcher/healthstyle/nutrition/internal/api"
	nutritionService "github.com/gerathewitcher/healthstyle/nutrition/internal/service/nutrition"
)

type serviceProvider struct {
	grpcConfig          config.GRPCConfig
	pgConfig            config.PGConfig
	dbClient            db.Client
	nutritionImpl       *nutrition.Implementation
	nutritionService    service.NutritionService
	nutritionRepository repository.NutritionRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()

		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {

	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create DB client: %v", err)
		}

		err = cl.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %s", err.Error())

		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) NutritionImpl(ctx context.Context) *nutrition.Implementation {

	if s.nutritionImpl == nil {

		s.nutritionImpl = nutrition.NewImplementation(s.NutritionService(ctx))
	}

	return s.nutritionImpl
}
func (s *serviceProvider) NutritionService(ctx context.Context) service.NutritionService {
	if s.nutritionService == nil {
		s.nutritionService = nutritionService.NewNutritionService(s.NutritionRepository(ctx))
	}
	return s.nutritionService
}
func (s *serviceProvider) NutritionRepository(ctx context.Context) repository.NutritionRepository {

	if s.nutritionRepository == nil {
		s.nutritionRepository = rep.NewNutritionRepository(s.DBClient(ctx))
	}
	return s.nutritionRepository
}
