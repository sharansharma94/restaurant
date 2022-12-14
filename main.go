package main

import (
	"fmt"
	"os"
	"restaurants-manager/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables from .env
	godotenv.Load()

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
	// routes.OrderItem(router)
	// routes.Invoice(router)

	fmt.Println("starting Server...")
	router.Run(":" + port)

}
