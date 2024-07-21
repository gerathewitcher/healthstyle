package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/converter"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i *Implementation) GetNutritionPlan(ctx context.Context, req *desc.GetNutritionPlanRequest) (*desc.GetNutritionPlanResponse, error) {

	plan, err := i.nutritionService.GetNutritionPlan(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &desc.GetNutritionPlanResponse{
		Plan: converter.ToNutritionPlanFromService(*plan),
	}, nil

}
