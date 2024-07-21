package nutrition

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/gerathewitcher/healthstyle/nutrition/internal/repository"
	"github.com/gerathewitcher/healthstyle/nutrition/pkg/db"
)

type repo struct {
	db db.Client
}

func NewNutritionRepository(db db.Client) repository.NutritionRepository {

	return &repo{db: db}
}

func (r *repo) getTotalRecords(ctx context.Context, tableName string) (uint64, error) {

	builder := sq.Select("count(*)").From(tableName)
	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	wrappedQuery := db.Query{
		Name:     "nutrition_repository.getTotalRecords",
		QueryRaw: query,
	}
	totalRows := r.db.DB().QueryRowContext(ctx, wrappedQuery, args...)
	var total uint64
	err = totalRows.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
