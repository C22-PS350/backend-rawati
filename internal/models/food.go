package models

type FoodResponse struct {
	FoodID   uint64 `json:"food_id"`
	Name     string `json:"name"`
	Calories uint64 `json:"calories"`
}
