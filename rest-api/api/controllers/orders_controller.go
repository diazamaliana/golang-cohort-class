package controllers

import (
    "github.com/gin-gonic/gin"
    "rest-api/api/models"
    "net/http"
)

// CreateOrder handles the creation of a new order
func CreateOrder(c *gin.Context) {
    var request models.Order
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // Create a new order in the database
    result := models.DB.Create(&request)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusOK, request)
}

// GetAllOrders retrieves all orders from the database
func GetAllOrders(c *gin.Context) {
    var orders []models.Order
    models.DB.Preload("Items").Find(&orders)
    c.JSON(http.StatusOK, orders)
}

// GetOrderById retrieves specific orders from the database
func GetOrderById(c *gin.Context) {
    id := c.Param("id") // Get the order ID from the URL parameter

    var order models.Order
    if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}
