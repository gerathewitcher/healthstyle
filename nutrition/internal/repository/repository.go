package repository

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

type NutritionRepository interface {
	GetNutritionPlan(ctx context.Context, id types.UUID) (*model.NutritionPlan, error)
	GetNutritionPlans(ctx context.Context, limit uint64, offset uint64) (*model.NutritionPlanList, error)
	CreateNutritionPlan(ctx context.Context, plan model.NutritionPlanToCreate) (types.UUID, error)
	GetFoods(ctx context.Context, limit uint64, offset uint64) (*model.FoodList, error)
	CreateFood(ctx context.Context, food model.FoodToCreate) (types.UUID, error)
	UpdateFood(ctx context.Context, food model.FoodToUpdate) error
	CreateMeal(ctx context.Context, meal model.MealToCreate) (types.UUID, error)
	CreateMealFood(ctx context.Context, mealFood model.MealFoodToCreate) (types.UUID, error)
}
