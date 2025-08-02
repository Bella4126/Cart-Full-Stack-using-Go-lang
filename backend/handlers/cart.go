package handlers

import (
	"net/http"
	"shopping-cart/database"
	"shopping-cart/models"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	var cartData struct {
		ItemID   uint `json:"item_id"`
		Quantity int  `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&cartData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if cartData.Quantity <= 0 {
		cartData.Quantity = 1
	}

	var existingCart models.Cart
	if err := database.DB.Where("user_id = ? AND item_id = ?", userID, cartData.ItemID).First(&existingCart).Error; err == nil {
		existingCart.Quantity += cartData.Quantity
		database.DB.Save(&existingCart)
		c.JSON(http.StatusOK, existingCart)
		return
	}

	cart := models.Cart{
		UserID:   userID,
		ItemID:   cartData.ItemID,
		Quantity: cartData.Quantity,
	}

	if err := database.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	database.DB.Preload("Item").First(&cart, cart.ID)
	c.JSON(http.StatusCreated, cart)
}

func GetCarts(c *gin.Context) {
	var carts []models.Cart
	database.DB.Preload("User").Preload("Item").Find(&carts)
	c.JSON(http.StatusOK, carts)
}