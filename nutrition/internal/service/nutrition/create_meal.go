package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (s *serv) CreateMeal(ctx context.Context, meal model.MealToCreate) (types.UUID, error) {
	return s.nutritionRepository.CreateMeal(ctx, meal)
}
