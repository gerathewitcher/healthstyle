package nutrition

import (
	"github.com/gerathewitcher/healthstyle/nutrition/internal/repository"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/service"
)

type serv struct {
	nutritionRepository repository.NutritionRepository
}

func NewNutritionService(nutritionRepository repository.NutritionRepository) service.NutritionService {
	return &serv{nutritionRepository: nutritionRepository}
}
