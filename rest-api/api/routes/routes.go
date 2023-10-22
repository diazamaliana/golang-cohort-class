package routes

import (
    "github.com/gin-gonic/gin"
    "rest-api/api/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    router.POST("/orders", controllers.CreateOrder)
    router.GET("/orders", controllers.GetAllOrders)
    router.GET("/order/:id", controllers.GetOrderById)
    router.PUT("/orders/:id", controllers.UpdateOrder)
    router.DELETE("/orders/:id", controllers.DeleteOrder)

    return router
}
