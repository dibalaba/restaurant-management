package routes

import (
	"github/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)

func Foodroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET()
	incomingRoutes.POST()
	incomingRoutes.PATCH
}
