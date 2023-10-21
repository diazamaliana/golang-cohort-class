package models

import (
    "gorm.io/gorm"
)

type Item struct {
    gorm.Model
    OrderID     uint     // Foreign key column
    Name        string
    Description string
    Quantity    int
}
