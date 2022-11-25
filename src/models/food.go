package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Name       string             `json:"name" validate:"required,min=2,max=100"`
	Price      float64            `json:"price" validate:"required"`
	Food_image string             `json:"food_image" validate:"required"`
	FoodID     string             `json:"food_id" validate:"required"`
	MenuID     string             `json:"menu_id" validate:"required"`
}

type FoodRequest struct {
	Name       string  `json:"name" validate:"required,min=2,max=100"`
	Price      float64 `json:"price" validate:"required"`
	Food_image string  `json:"food_image" validate:"required"`
	FoodID     string  `json:"food_id" validate:"required"`
	MenuID     string  `json:"menu_id" validate:"required"`
}
