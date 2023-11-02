package dto

type Customer struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    CustomerID string `json:"customerID" validate:"required"`
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Phone string `json:"phone" validate:"required"`
    Address string `json:"address" validate:"required"`
    Age int `json:"age" validate:"required"`
    }

