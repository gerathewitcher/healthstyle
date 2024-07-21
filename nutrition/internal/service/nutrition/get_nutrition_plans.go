package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
)

func (s *serv) GetNutritionPlans(ctx context.Context, limit uint64, offset uint64) (*model.NutritionPlanList, error) {
	plans, err := s.nutritionRepository.GetNutritionPlans(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return plans, nil
}
