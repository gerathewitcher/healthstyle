package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i *Implementation) CreateFood(ctx context.Context, req *desc.CreateFoodRequest) (*desc.CreateFoodResponse, error) {
	foodToCreate := model.FoodToCreate{
		Name: req.Name,
	}

	if req.Calorie != nil {
		foodToCreate.Calorie = &req.Calorie.Value
	}

	if req.Carbs != nil {
		foodToCreate.Carbs = &req.Carbs.Value
	}

	if req.Proteins != nil {
		foodToCreate.Proteins = &req.Proteins.Value
	}

	if req.Fats != nil {
		foodToCreate.Fats = &req.Fats.Value
	}

	foodId, err := i.nutritionService.CreateFood(ctx, foodToCreate)

	if err != nil {
		return nil, err
	}

	return &desc.CreateFoodResponse{Id: foodId}, nil
}
