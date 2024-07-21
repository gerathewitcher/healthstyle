package nutrition

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (r *repo) CreateMeal(ctx context.Context, mealData model.MealToCreate) (types.UUID, error) {

	builder := sq.Insert("meal").
		PlaceholderFormat(sq.Dollar).
		Columns(
			"nutrition_plan_id",
			"name",
			"time",
		).
		Values(mealData.NutritionPlanId, mealData.Name, mealData.Time).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return "", err
	}

	wrappedQuery := db.Query{Name: "nutrition_repository.CreateMeal", QueryRaw: query}
	var mealId types.UUID

	err = r.db.DB().QueryRowContext(ctx, wrappedQuery, args...).Scan(&mealId)

	if err != nil {
		return "", err
	}

	return mealId, nil
}
