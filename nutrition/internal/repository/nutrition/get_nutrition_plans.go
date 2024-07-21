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

func (r *repo) GetNutritionPlans(ctx context.Context, limit uint64, offset uint64) (*model.NutritionPlanList, error) {

	totalRecords, err := r.getTotalRecords(ctx, "nutrition_plan")

	if err != nil {
		return nil, err
	}

	builder := sq.Select(
		"id",
		"day",
		"name",
		"created_at",
		"updated_at",
	).
		From("nutrition_plan").Offset(offset).Limit(limit)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("error while generating query %w", err)
	}

	wrappedQuery := db.Query{
		Name:     "nutrition_repository.GetNutritionPlans",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, wrappedQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error while querying plans %w", err)
	}

	defer rows.Close()

	plans := []modelRepo.NutritionPlanShort{}
	planList := modelRepo.NutritionPlanList{
		Plans: plans,
		Total: totalRecords,
	}

	for rows.Next() {
		var plan modelRepo.NutritionPlanShort
		err := rows.Scan(
			&plan.ID,
			&plan.Day,
			&plan.Name,
			&plan.CreatedAt,
			&plan.UpdatedAt,
		)

		planList.Plans = append(planList.Plans, plan)

		if err != nil {
			return nil, fmt.Errorf("error while processing plans rows %w", err)
		}

	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error while retrieving plans %w", rows.Err())
	}

	convertedPlanList := converter.ToNutritionPlanListFromRepo(planList)

	return &convertedPlanList, nil
}
