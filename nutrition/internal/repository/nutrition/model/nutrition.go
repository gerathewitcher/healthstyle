package model

import (
	"database/sql"
	"time"

	"github.com/gerathewitcher/healthstyle/nutrition/pkg/types"
)

type NutritionPlanList struct {
	Total uint64
	Plans []NutritionPlanShort
}

type FoodList struct {
	Total uint64
	Foods []Food
}

type NutritionPlanToCreate struct {
	Name string
	Day  time.Time
}

type NutritionPlanShort struct {
	ID        types.UUID
	Name      string
	Day       time.Time
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NutritionPlan struct {
	ID        types.UUID
	Name      string
	Day       time.Time
	Meals     []Meal
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Food struct {
	ID        types.UUID
	Name      string
	Calorie   sql.NullInt32
	Proteins  sql.NullInt32
	Fats      sql.NullInt32
	Carbs     sql.NullInt32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Meal struct {
	ID        types.UUID
	Name      string
	Time      time.Time
	CreatedAt time.Time
	UpdatedAt sql.NullTime
	Foods     []MealFood
}

// MealFood representing food which is related with meal
type MealFood struct {
	ID        types.UUID
	Food      Food
	Weight    sql.NullInt32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
