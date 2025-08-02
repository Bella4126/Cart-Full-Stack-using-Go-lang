package main

import (
	"os"
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
		AllowOrigins:     []string{"http://localhost:5173", "https://your-frontend-domain.com"}, // Add your frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Add a root route for health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Cart API is running!",
			"status":  "healthy",
		})
	})

	// Add an API health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"api":    "cart-api",
		})
	})

	// Your existing routes
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

	// Use PORT environment variable from Render, fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}