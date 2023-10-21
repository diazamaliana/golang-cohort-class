package models

import (
    "gorm.io/gorm"
    "time"
)

type Order struct {
    gorm.Model
    CustomerName string
    OrderedAt    time.Time
    Items        []Item // Relationship with the Item model
}
