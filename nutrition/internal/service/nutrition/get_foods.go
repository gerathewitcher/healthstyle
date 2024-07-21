package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
)

func (s *serv) GetFoods(ctx context.Context, limit uint64, offset uint64) (*model.FoodList, error) {
	foodList, err := s.nutritionRepository.GetFoods(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return foodList, nil
}
