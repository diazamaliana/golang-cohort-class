package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(connectionURL string) {
    db, err := gorm.Open(mysql.Open(connectionURL), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    DB = db
    fmt.Println("Successfully connected to database")
}
