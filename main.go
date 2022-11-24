package main

import (
	"fmt"
	"os"
	"restaurants-manager/routes"

	"github.com/gin-gonic/gin"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())

	routes.Food(router)
	// routes.Menu(router)
	// routes.Table(router)
	// routes.Order(router)
	routes.OrderItem(router)
	// routes.Invoice(router)

	fmt.Println("starting Server...")
	router.Run(":" + port)

}
