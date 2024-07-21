package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (s *serv) GetNutritionPlan(ctx context.Context, id types.UUID) (*model.NutritionPlan, error) {
	plan, err := s.nutritionRepository.GetNutritionPlan(ctx, id)
	if err != nil {
		return nil, err
	}
	return plan, nil
}
