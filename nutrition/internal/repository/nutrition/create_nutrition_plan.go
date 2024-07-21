package nutrition

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (r *repo) CreateNutritionPlan(ctx context.Context, plan model.NutritionPlanToCreate) (types.UUID, error) {
	builder := sq.Insert("nutrition_plan").
		PlaceholderFormat(sq.Dollar).
		Columns(
			"name",
			"day").
		Values(
			plan.Name,
			plan.Day,
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()

	if err != nil {
		return "", err
	}

	wrappedQuery := db.Query{Name: "nutrition_repository.CreateNutritionPlan", QueryRaw: query}

	var planId types.UUID

	err = r.db.DB().QueryRowContext(ctx, wrappedQuery, args...).Scan(&planId)

	if err != nil {
		return "", err
	}

	return planId, nil
}
