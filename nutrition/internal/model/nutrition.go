package model

import (
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
	UpdatedAt *time.Time
}

type NutritionPlan struct {
	ID        types.UUID
	Name      string
	Day       time.Time
	Meals     []Meal
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Food struct {
	ID        types.UUID
	Name      string
	Calorie   *uint32
	Proteins  *uint32
	Fats      *uint32
	Carbs     *uint32
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type FoodToCreate struct {
	Name     string
	Calorie  *uint32
	Proteins *uint32
	Fats     *uint32
	Carbs    *uint32
}

type FoodToUpdate struct {
	ID       types.UUID
	Name     *string
	Calorie  *uint32
	Proteins *uint32
	Fats     *uint32
	Carbs    *uint32
}

type Meal struct {
	ID        types.UUID
	Name      string
	Time      time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
	Foods     []MealFood
}

type MealToCreate struct {
	NutritionPlanId types.UUID
	Name            string
	Time            time.Time
}

// MealFood representing food which is related with meal
type MealFood struct {
	ID        types.UUID
	Food      Food
	Weight    *uint32
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type MealFoodToCreate struct {
	FoodId types.UUID
	MealId types.UUID
	Weight *uint32
}
