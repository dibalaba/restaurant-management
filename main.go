package main

import (
	"github/gin-gonic/gin"
	"golang-restaurant-management/controllers"
	"golang-restaurant-management/database"
	"golang-restaurant-management/helpers"
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/models"
	"golang-restaurant-management/routes"
	"os"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {

		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}
