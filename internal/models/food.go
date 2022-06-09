package models

type FoodResponse struct {
	FoodID   uint64  `json:"food_id"`
	Name     string  `json:"name"`
	Calories float64 `json:"calories"`
}

type FoodTest1 struct {
	FoodID   uint64  `faker:"unique"`
	Name     string  `faker:"-"`
	Calories float64 `faker:"-"`
}
