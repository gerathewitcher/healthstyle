package nutrition

import (
	"github.com/gerathewitcher/healthstyle/nutrition/internal/service"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

type Implementation struct {
	desc.UnimplementedNutritionV1Server
	nutritionService service.NutritionService
}

func NewImplementation(nutritionService service.NutritionService) *Implementation {

	return &Implementation{nutritionService: nutritionService}
}
