package handlers

import (
	"net/http"
	"shopping-cart/database"
	"shopping-cart/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Select("id, username, created_at, updated_at").Find(&users)
	c.JSON(http.StatusOK, users)
}