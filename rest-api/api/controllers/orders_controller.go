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

// UpdateOrder updates an existing order
func UpdateOrder(c *gin.Context) {
    id := c.Param("id")

    // Retrieve the order from the database, including associated items
    var order models.Order
    if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    // Bind the updated order and items from the request body
    var updatedOrder models.Order
    if err := c.ShouldBindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the items provided in the request have the correct order ID
    for _, updatedItem := range updatedOrder.Items {
        if updatedItem.OrderID != order.ID {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Item does not belong to the order"})
            return
        }
    }

    // Update the order details
    models.DB.Model(&order).Updates(updatedOrder)

    // Create a map to keep track of items to update or create
    itemUpdates := make(map[uint]models.Item)
    for _, updatedItem := range updatedOrder.Items {
        // If an item has an ID, it should be updated; otherwise, it's new
        if updatedItem.ID != 0 {
            itemUpdates[updatedItem.ID] = updatedItem
        }
    }

    // Update or create the associated items
    for i, item := range order.Items {
        if updatedItem, exists := itemUpdates[item.ID]; exists {
            // Update existing item with data from updatedItem
            models.DB.Model(&item).Updates(updatedItem)
            // Remove the updated item from the map
            delete(itemUpdates, item.ID)
        } else {
            // Delete any items not included in the updated request
            models.DB.Delete(&item)
            // Remove the deleted item from the order's item slice
            order.Items = append(order.Items[:i], order.Items[i+1:]...)
        }
    }

    // Create new items for any remaining entries in itemUpdates
    for _, newItem := range itemUpdates {
        order.Items = append(order.Items, newItem)
    }

    // Save the updated order with items
    models.DB.Save(&order)

    c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order by ID
func DeleteOrder(c *gin.Context) {
    id := c.Param("id")

    // Retrieve the order from the database
    var order models.Order
    if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    // Delete the associated items
    for _, item := range order.Items {
        models.DB.Delete(&item)
    }

    // Delete the order itself
    models.DB.Delete(&order)

    c.JSON(http.StatusOK, gin.H{"message": "Order and associated items deleted"})
}

