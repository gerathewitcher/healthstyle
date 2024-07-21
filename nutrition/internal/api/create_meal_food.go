package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i Implementation) CreateMealFood(ctx context.Context, mealFood *desc.CreateMealFoodRequest) (*desc.CreateMealFoodResponse, error) {
	mealFoodToCreate := model.MealFoodToCreate{
		MealId: mealFood.MealId,
		FoodId: mealFood.FoodId}

	if mealFood.FoodWeight != nil {
		mealFoodToCreate.Weight = &mealFood.FoodWeight.Value
	}
	mealFoodId, err := i.nutritionService.CreateMealFood(ctx,
		mealFoodToCreate)

	if err != nil {
		return nil, err
	}

	return &desc.CreateMealFoodResponse{
		Id: mealFoodId,
	}, err

}
