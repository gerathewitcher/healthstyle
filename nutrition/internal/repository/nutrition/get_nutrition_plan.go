package nutrition

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition/converter"
	modelRepo "github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

func (r *repo) GetNutritionPlan(ctx context.Context, id types.UUID) (*model.NutritionPlan, error) {

	queryName := "nutrition_repository.GetNutritionPlan"
	builder := sq.Select(
		"np.id",
		"np.day",
		"np.name",
		"np.created_at",
		"np.updated_at",
		"m.id",
		"m.name",
		"m.time",
		"m.created_at",
		"m.updated_at",
		"mf.id",
		"f.id",
		"f.name",
		"f.calorie",
		"f.proteins",
		"f.fats",
		"f.carbs",
		"mf.weight",
	).PlaceholderFormat(sq.Dollar).
		From("nutrition_plan as np").
		LeftJoin("meal as m on m.nutrition_plan_id = np.id").
		LeftJoin("meal_food as mf on mf.meal_id = m.id").
		LeftJoin("food as f  on mf.food_id = f.id").
		Where(sq.Eq{"np.id": id})

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, fmt.Errorf("error while generating query %w", err)
	}

	wrappedQuery := db.Query{
		Name:     queryName,
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, wrappedQuery, args...)

	if err != nil {
		return nil, fmt.Errorf("error while querying nutrition plan %w", err)
	}
	defer rows.Close()
	var plan modelRepo.NutritionPlan
	meals := make(map[types.UUID]*modelRepo.Meal)

	for rows.Next() {
		var mealID sql.NullString
		var mealTime sql.NullTime
		var mealName sql.NullString
		var mealCreatedAt sql.NullTime
		var mealUpdatedAt sql.NullTime
		var mealFoodID sql.NullString
		var foodID sql.NullString
		var foodName sql.NullString
		var foodWeight sql.NullInt32
		var foodCalorie sql.NullInt32
		var foodProteins sql.NullInt32
		var foodFats sql.NullInt32
		var foodCarbs sql.NullInt32

		err := rows.Scan(
			&plan.ID,
			&plan.Day,
			&plan.Name,
			&plan.CreatedAt,
			&plan.UpdatedAt,
			&mealID,
			&mealName,
			&mealTime,
			&mealCreatedAt,
			&mealUpdatedAt,
			&mealFoodID,
			&foodID,
			&foodName,
			&foodCalorie,
			&foodProteins,
			&foodFats,
			&foodCarbs,
			&foodWeight,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning nutrition plan rows %w", err)
		}

		mealFoodFromSql := func() modelRepo.MealFood {
			mealFood := modelRepo.MealFood{
				ID:     mealFoodID.String,
				Weight: foodWeight,
			}

			food := modelRepo.Food{
				ID:       foodID.String,
				Name:     foodName.String,
				Calorie:  foodCalorie,
				Proteins: foodProteins,
				Fats:     foodFats,
				Carbs:    foodCarbs,
			}

			mealFood.Food = food

			return mealFood
		}

		if mealID.Valid {
			meal, ok := meals[mealID.String]

			if !ok {
				meal = &modelRepo.Meal{
					ID:        mealID.String,
					Name:      mealName.String,
					Time:      mealTime.Time,
					CreatedAt: mealCreatedAt.Time,
					UpdatedAt: mealUpdatedAt,
				}

				meals[mealID.String] = meal

			}
			if mealFoodID.Valid {
				mealFood := mealFoodFromSql()
				meal.Foods = append(meal.Foods, mealFood)

			}
		}

	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error while retrieving plans %w", rows.Err())
	}
	for _, meal := range meals {

		plan.Meals = append(plan.Meals, *meal)
	}

	convertedPlan := converter.ToNutritionPlanFromRepo(&plan)

	return &convertedPlan, nil
}
