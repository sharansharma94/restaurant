package routes

import (
	controller "restaurants-manager/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItem(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems-order/:id", controller.GetOrderItemsByOrder())

}
