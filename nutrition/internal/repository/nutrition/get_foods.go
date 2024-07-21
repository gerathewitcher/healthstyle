package nutrition

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition/converter"
	modelRepo "github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
)

func (r *repo) GetFoods(ctx context.Context, limit uint64, offset uint64) (*model.FoodList, error) {
	totalRecords, err := r.getTotalRecords(ctx, "food")
	if err != nil {
		return nil, fmt.Errorf("error while getting total food records %w", err)
	}
	builder := sq.Select("id",
		"name",
		"calorie",
		"carbs",
		"proteins",
		"fats",
		"created_at",
		"updated_at").From("food").Limit(limit).Offset(offset)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error while generating query %w", err)
	}

	wrappedQuery := db.Query{
		Name:     "nutrition_repository.GetFoods",
		QueryRaw: query,
	}
	rows, err := r.db.DB().QueryContext(ctx, wrappedQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error while querying foods %w", err)
	}
	defer rows.Close()
	foods := []modelRepo.Food{}
	foodList := modelRepo.FoodList{
		Total: totalRecords,
		Foods: foods,
	}

	for rows.Next() {
		var food modelRepo.Food
		err := rows.Scan(
			&food.ID,
			&food.Name,
			&food.Calorie,
			&food.Carbs,
			&food.Proteins,
			&food.Fats,
			&food.CreatedAt,
			&food.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error while scanning food %w", err)
		}
		foodList.Foods = append(foodList.Foods, food)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error while scanning food %w", err)
	}

	convertedFoodList := converter.ToFoodListFromRepo(foodList)
	return &convertedFoodList, nil
}
