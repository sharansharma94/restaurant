package routes

import (
	controller "restaurants-manager/controllers"

	"github.com/gin-gonic/gin"
)

func Food(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food/", controller.GetFoods())
	incomingRoutes.GET("/food/:id", controller.GetFood())
	incomingRoutes.POST("/food", controller.CreateFood())
	incomingRoutes.PATCH("/food/:id", controller.UpdateFood())
}
