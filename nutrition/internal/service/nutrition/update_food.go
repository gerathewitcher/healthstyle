package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
)

func (s *serv) UpdateFood(ctx context.Context, food model.FoodToUpdate) error {
	err := s.nutritionRepository.UpdateFood(ctx, food)
	if err != nil {
		return err
	}
	return nil
}
