package handlers

import (
	"net/http"
	"shopping-cart/database"
	"shopping-cart/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var cartItems []models.Cart
	if err := database.DB.Where("user_id = ?", userID).Preload("Item").Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	var total float64
	for _, cart := range cartItems {
		total += cart.Item.Price * float64(cart.Quantity)
	}

	order := models.Order{
		UserID: userID,
		Total:  total,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	for _, cart := range cartItems {
		orderItem := models.OrderItem{
			OrderID:  order.ID,
			ItemID:   cart.ItemID,
			Quantity: cart.Quantity,
			Price:    cart.Item.Price,
		}
		database.DB.Create(&orderItem)
	}

	database.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	database.DB.Preload("User").First(&order, order.ID)
	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("User").Find(&orders)
	c.JSON(http.StatusOK, orders)
}