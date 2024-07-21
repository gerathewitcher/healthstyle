package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/converter"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i *Implementation) GetNutritionPlans(ctx context.Context, req *desc.GetNutritionPlansRequest) (*desc.GetNutritionPlansResponse, error) {
	planList, err := i.nutritionService.GetNutritionPlans(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &desc.GetNutritionPlansResponse{
		Total: planList.Total,
		Plans: converter.ToNutritionPlansFromService(planList.Plans),
	}, nil
}
