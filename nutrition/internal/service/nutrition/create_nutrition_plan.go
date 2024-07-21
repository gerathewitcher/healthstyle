package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (s *serv) CreateNutritionPlan(ctx context.Context, plan model.NutritionPlanToCreate) (types.UUID, error) {
	return s.nutritionRepository.CreateNutritionPlan(ctx, plan)
}
