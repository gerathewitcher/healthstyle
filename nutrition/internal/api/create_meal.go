package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i Implementation) CreateMeal(ctx context.Context, req *desc.CreateMealRequest) (*desc.CreateMealResponse, error) {
	mealId, err := i.nutritionService.CreateMeal(ctx, model.MealToCreate{
		NutritionPlanId: req.PlanId,
		Name:            req.Name,
		Time:            req.MealTime.AsTime(),
	})

	if err != nil {
		return nil, err
	}
	return &desc.CreateMealResponse{Id: mealId}, nil
}
