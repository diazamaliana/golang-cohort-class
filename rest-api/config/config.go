package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func GetDBConfig() string {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")

    // Construct the database connection string
    dbConfig := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"

    return dbConfig
}
