package nutrition

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/errors"
)

func (r *repo) UpdateFood(ctx context.Context, food model.FoodToUpdate) error {
	builder := sq.Update("food").
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": food.ID}).Set("updated_at", sq.Expr("now()"))

	if food.Name != nil {
		builder = builder.Set("name", *food.Name)
	}
	if food.Calorie != nil {
		builder = builder.Set("calorie", *food.Calorie)
	}
	if food.Proteins != nil {
		builder = builder.Set("proteins", *food.Proteins)
	}
	if food.Fats != nil {
		builder = builder.Set("fats", *food.Fats)
	}
	if food.Carbs != nil {
		builder = builder.Set("carbs", *food.Carbs)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	wrappedQuery := db.Query{Name: "nutrition_repository.UpdateFood", QueryRaw: query}

	commandTag, err := r.db.DB().ExecContext(ctx, wrappedQuery, args...)
	if commandTag.RowsAffected() == 0 {
		return errors.RecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}
