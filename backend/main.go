package main

import (
	"shopping-cart/database"
	"shopping-cart/handlers"
	"shopping-cart/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	r.POST("/users", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.POST("/users/login", handlers.Login)
	r.POST("/items", handlers.CreateItem)
	r.GET("/items", handlers.GetItems)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/carts", handlers.AddToCart)
		protected.GET("/carts", handlers.GetCarts)
		protected.POST("/orders", handlers.CreateOrder)
		protected.GET("/orders", handlers.GetOrders)
	}

	r.Run(":8080")
}