package nutrition

import (
	"context"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/converter"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/errors"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) UpdateFood(ctx context.Context, food *desc.UpdateFoodRequest) (*emptypb.Empty, error) {

	err := i.nutritionService.UpdateFood(ctx, converter.ToFoodToUpdateFromDesc(food))
	if err != nil {
		if errors.Is(err, errors.RecordNotFound) {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
