package controller

import (
	"context"
	"fmt"
	"net/http"
	"restaurants-manager/src/database"
	"restaurants-manager/src/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		foodId := c.Param("id")
		var food models.Food

		var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
		err := foodCollection.FindOne(ctx, bson.M{"id": foodId}).Decode(&food)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, food)
	}
}

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// var foods []models.Food

		var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

		cur, err := foodCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var results []*models.Food
		for cur.Next(ctx) {
			var food models.Food

			if err := cur.Decode(&food); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			fmt.Println(food.Price)
			results = append(results, &food)

		}
		c.JSON(http.StatusOK, results)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		var food models.FoodRequest

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("fooddata : ", food)

		var collection *mongo.Collection = database.OpenCollection(database.Client, "food")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := collection.InsertOne(ctx, &food)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, result)

	}

}
