package main

import (
    "rest-api/config"
    "rest-api/api/models"
    "rest-api/api/routes"
)

func main() {
    // Load environment variables from .env
    config.LoadEnv()

    // Get the database connection string from environment variables
    dbConfig := config.GetDBConfig()

    // Connect to the database using the config
    models.ConnectDB(dbConfig)

    // Initialize the Gin router
    router := routes.SetupRouter()

    // Start the server
    router.Run(":8080")
}
