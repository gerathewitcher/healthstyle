package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (s *serv) CreateFood(ctx context.Context, food model.FoodToCreate) (types.UUID, error) {
	return s.nutritionRepository.CreateFood(ctx, food)
}
