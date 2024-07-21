package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (s *serv) CreateMealFood(ctx context.Context, mealFood model.MealFoodToCreate) (types.UUID, error) {

	return s.nutritionRepository.CreateMealFood(ctx, mealFood)
}
