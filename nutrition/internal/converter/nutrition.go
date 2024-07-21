package converter

import (
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	desc "github.com/gerathewitcher/healthstyle/nutrition/pkg/nutrition_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNutritionPlansFromService(plans []model.NutritionPlanShort) []*desc.NutritionPlanShort {
	convertedPlans := make([]*desc.NutritionPlanShort, len(plans))
	for i, plan := range plans {
		convertedPlans[i] = ToNutritionPlanShortFromService(plan)
	}
	return convertedPlans
}

func ToFoodsFromService(foods []model.Food) []*desc.Food {
	convertedFoods := make([]*desc.Food, len(foods))
	for i, food := range foods {
		convertedFoods[i] = ToFoodFromService(food)
	}
	return convertedFoods
}

func ToMealFromService(meal model.Meal) *desc.Meal {
	mealFoods := make([]*desc.MealFood, len(meal.Foods))
	for i, mealFood := range meal.Foods {
		mealFoods[i] = toMealFoodFromService(mealFood)
	}

	convertedMeal := &desc.Meal{
		Id:        meal.ID,
		Name:      meal.Name,
		MealFoods: mealFoods,
		Time:      timestamppb.New(meal.Time),
		CreatedAt: timestamppb.New(meal.CreatedAt),
	}
	if meal.UpdatedAt != nil {
		convertedMeal.UpdatedAt = timestamppb.New(*meal.UpdatedAt)
	}
	return convertedMeal
}

func toMealFoodFromService(mealFood model.MealFood) *desc.MealFood {
	convertedMealFood := &desc.MealFood{
		Id:   mealFood.ID,
		Food: ToFoodFromService(mealFood.Food),

		CreatedAt: timestamppb.New(mealFood.CreatedAt),
	}

	if mealFood.UpdatedAt != nil {
		convertedMealFood.UpdatedAt = timestamppb.New(*mealFood.UpdatedAt)
	}

	if mealFood.Weight != nil {
		convertedMealFood.Weight = wrapperspb.UInt32(*mealFood.Weight)
	}
	return convertedMealFood
}

func ToFoodFromService(food model.Food) *desc.Food {
	convertedFood := &desc.Food{
		Id:        food.ID,
		Name:      food.Name,
		CreatedAt: timestamppb.New(food.CreatedAt),
	}

	if food.UpdatedAt != nil {
		convertedFood.UpdatedAt = timestamppb.New(*food.UpdatedAt)
	}

	if food.Calorie != nil {
		convertedFood.Calorie = wrapperspb.UInt32(*food.Calorie)
	}
	if food.Carbs != nil {
		convertedFood.Carbs = wrapperspb.UInt32(*food.Carbs)
	}
	if food.Fats != nil {
		convertedFood.Fats = wrapperspb.UInt32(*food.Fats)
	}
	if food.Proteins != nil {
		convertedFood.Proteins = wrapperspb.UInt32(*food.Proteins)
	}
	return convertedFood
}

func ToNutritionPlanFromService(plan model.NutritionPlan) *desc.NutritionPlan {
	convertedPlan := &desc.NutritionPlan{
		Name:      plan.Name,
		Id:        plan.ID,
		Day:       timestamppb.New(plan.Day),
		CreatedAt: timestamppb.New(plan.CreatedAt),
	}
	if plan.UpdatedAt != nil {
		convertedPlan.UpdatedAt = timestamppb.New(*plan.UpdatedAt)
	}

	meals := make([]*desc.Meal, len(plan.Meals))

	for i, meal := range plan.Meals {
		meals[i] = ToMealFromService(meal)
	}

	convertedPlan.Meals = meals

	return convertedPlan
}

func ToNutritionPlanShortFromService(plan model.NutritionPlanShort) *desc.NutritionPlanShort {
	convertedPlan := &desc.NutritionPlanShort{
		Name:      plan.Name,
		Id:        plan.ID,
		Day:       timestamppb.New(plan.Day),
		CreatedAt: timestamppb.New(plan.CreatedAt),
	}

	if plan.UpdatedAt != nil {
		convertedPlan.UpdatedAt = timestamppb.New(*plan.UpdatedAt)
	}

	return convertedPlan
}

func ToFoodToUpdateFromDesc(food *desc.UpdateFoodRequest) model.FoodToUpdate {
	foodToUpdate := model.FoodToUpdate{
		ID: food.Id,
	}

	if food.Calorie != nil {
		foodToUpdate.Calorie = &food.Calorie.Value
	}

	if food.Carbs != nil {
		foodToUpdate.Carbs = &food.Carbs.Value
	}

	if food.Fats != nil {
		foodToUpdate.Fats = &food.Fats.Value
	}

	if food.Proteins != nil {
		foodToUpdate.Proteins = &food.Proteins.Value
	}

	if food.Name != nil {
		foodToUpdate.Name = &food.Name.Value
	}
	return foodToUpdate
}
