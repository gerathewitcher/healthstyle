package nutrition

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (r *repo) CreateFood(ctx context.Context, food model.FoodToCreate) (types.UUID, error) {
	builder := sq.Insert("food").Columns(
		"name",
		"calorie",
		"proteins",
		"fats",
		"carbs").
		Values(
			food.Name,
			food.Calorie,
			food.Proteins,
			food.Fats,
			food.Carbs).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return "", err
	}

	wrappedQuery := db.Query{Name: "nutrition_repository.CreateFood", QueryRaw: query}
	var foodId types.UUID
	err = r.db.DB().QueryRowContext(ctx, wrappedQuery, args...).Scan(&foodId)

	if err != nil {
		return "", err
	}
	return foodId, nil
}
