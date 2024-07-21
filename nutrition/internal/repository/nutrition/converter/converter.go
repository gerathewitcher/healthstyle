package converter

import (
	"github.com/gerathewitcher/healthstyle/nutrition/internal/model"
	modelRepo "github.com/gerathewitcher/healthstyle/nutrition/internal/repository/nutrition/model"
)

func ToFoodFromRepo(repoFood modelRepo.Food) model.Food {
	convertedFood := model.Food{
		ID:        repoFood.ID,
		Name:      repoFood.Name,
		CreatedAt: repoFood.CreatedAt,
	}

	if repoFood.Calorie.Valid {
		caloire := uint32(repoFood.Calorie.Int32)
		convertedFood.Calorie = &caloire
	}
	if repoFood.Proteins.Valid {
		proteins := uint32(repoFood.Proteins.Int32)
		convertedFood.Proteins = &proteins
	}
	if repoFood.Fats.Valid {
		fats := uint32(repoFood.Fats.Int32)
		convertedFood.Fats = &fats
	}
	if repoFood.Carbs.Valid {
		carbs := uint32(repoFood.Carbs.Int32)
		convertedFood.Carbs = &carbs
	}
	if repoFood.UpdatedAt.Valid {
		convertedFood.UpdatedAt = &repoFood.UpdatedAt.Time
	}
	return convertedFood
}

func ToMealFoodFromRepo(repoMealFood modelRepo.MealFood) model.MealFood {
	convertedMealFood := model.MealFood{
		ID:        repoMealFood.ID,
		Food:      ToFoodFromRepo(repoMealFood.Food),
		CreatedAt: repoMealFood.CreatedAt,
	}

	if repoMealFood.Weight.Valid {
		weigth := uint32(repoMealFood.Weight.Int32)
		convertedMealFood.Weight = &weigth
	}
	if repoMealFood.UpdatedAt.Valid {
		convertedMealFood.UpdatedAt = &repoMealFood.UpdatedAt.Time
	}

	return convertedMealFood
}

func ToMealFromRepo(repoMeal modelRepo.Meal) model.Meal {
	convertedMeal := model.Meal{
		ID:        repoMeal.ID,
		Name:      repoMeal.Name,
		Time:      repoMeal.Time,
		CreatedAt: repoMeal.CreatedAt,
	}

	convertedMealFoods := make([]model.MealFood, len(repoMeal.Foods))

	if repoMeal.UpdatedAt.Valid {
		convertedMeal.UpdatedAt = &repoMeal.UpdatedAt.Time
	}

	for i, mealFood := range repoMeal.Foods {
		convertedMealFoods[i] = ToMealFoodFromRepo(mealFood)
	}
	convertedMeal.Foods = convertedMealFoods

	return convertedMeal
}

func ToNutritionPlanFromRepo(repoNutritionPlan *modelRepo.NutritionPlan) model.NutritionPlan {
	convertedPlan := model.NutritionPlan{
		ID:        repoNutritionPlan.ID,
		Name:      repoNutritionPlan.Name,
		Day:       repoNutritionPlan.Day,
		CreatedAt: repoNutritionPlan.CreatedAt,
	}
	convertedMeals := make([]model.Meal, len(repoNutritionPlan.Meals))

	for i, meal := range repoNutritionPlan.Meals {
		convertedMeals[i] = ToMealFromRepo(meal)
	}
	if repoNutritionPlan.UpdatedAt.Valid {
		convertedPlan.UpdatedAt = &repoNutritionPlan.UpdatedAt.Time
	}
	convertedPlan.Meals = convertedMeals

	return convertedPlan
}

func ToNutritionPlanShortFromRepo(repoNutritionPlan modelRepo.NutritionPlanShort) model.NutritionPlanShort {
	convertedPlan := model.NutritionPlanShort{
		ID:        repoNutritionPlan.ID,
		Name:      repoNutritionPlan.Name,
		Day:       repoNutritionPlan.Day,
		CreatedAt: repoNutritionPlan.CreatedAt,
	}

	if repoNutritionPlan.UpdatedAt.Valid {
		convertedPlan.UpdatedAt = &repoNutritionPlan.UpdatedAt.Time
	}
	return convertedPlan
}

func ToNutritionPlanListFromRepo(repoNutritionPlanList modelRepo.NutritionPlanList) model.NutritionPlanList {
	convertedPlanList := model.NutritionPlanList{
		Total: repoNutritionPlanList.Total,
	}
	convertedPlans := make([]model.NutritionPlanShort, len(repoNutritionPlanList.Plans))

	for i, plan := range repoNutritionPlanList.Plans {
		convertedPlans[i] = ToNutritionPlanShortFromRepo(plan)
	}
	convertedPlanList.Plans = convertedPlans
	return convertedPlanList
}

func ToFoodListFromRepo(repoFoodList modelRepo.FoodList) model.FoodList {
	convertedFoodList := model.FoodList{
		Total: repoFoodList.Total,
	}
	convertedFoods := make([]model.Food, len(repoFoodList.Foods))

	for i, food := range repoFoodList.Foods {
		convertedFoods[i] = ToFoodFromRepo(food)
	}
	convertedFoodList.Foods = convertedFoods
	return convertedFoodList
}
