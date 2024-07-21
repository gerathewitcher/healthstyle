package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/converter"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
)

func (i *Implementation) GetFoods(ctx context.Context, req *desc.GetFoodsRequest) (*desc.GetFoodsResponse, error) {
	foodList, err := i.nutritionService.GetFoods(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	return &desc.GetFoodsResponse{
		Total: foodList.Total,
		Foods: converter.ToFoodsFromService(foodList.Foods),
	}, nil
}
