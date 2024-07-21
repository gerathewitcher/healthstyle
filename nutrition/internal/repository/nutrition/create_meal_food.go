package nutrition

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (r *repo) CreateMealFood(ctx context.Context, mealFood model.MealFoodToCreate) (types.UUID, error) {
	log.Println(mealFood.Weight)
	builder := sq.Insert("meal_food").PlaceholderFormat(sq.Dollar).Columns(
		"meal_id",
		"food_id",
		"weight",
	).Values(mealFood.MealId, mealFood.FoodId, mealFood.Weight).Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return "", err
	}

	wrappedQuery := db.Query{Name: "nutrition_repository.CreateMealFood", QueryRaw: query}

	var mealFoodId types.UUID

	err = r.db.DB().QueryRowContext(ctx, wrappedQuery, args...).Scan(&mealFoodId)

	if err != nil {
		return "", err
	}

	return mealFoodId, nil
}
