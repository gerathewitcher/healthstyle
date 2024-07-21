package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i *Implementation) CreateNutritionPlan(ctx context.Context, req *desc.CreateNutritionPlanRequest) (*desc.CreateNutritionPlanResponse, error) {
	planId, err := i.nutritionService.CreateNutritionPlan(ctx, model.NutritionPlanToCreate{
		Name: req.Name,
		Day:  req.Day.AsTime(),
	})

	if err != nil {
		return nil, err
	}

	return &desc.CreateNutritionPlanResponse{
		Id: planId,
	}, nil
}
