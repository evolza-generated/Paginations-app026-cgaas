package dto

type Driver struct {
    ObjectId string `gorm:"column:objectId" json:"objectId" validate:"required"`
    DriverId string `json:"driverId" validate:"required"`
    FirstName string `json:"firstName" validate:"required"`
    LastName string `json:"lastName" validate:"required"`
    Email string `json:"email" validate:"required"`
    Phone string `json:"phone" validate:"required"`
    LicenseNumber string `json:"licenseNumber" validate:"required"`
    LicenseExpirationDate string `json:"licenseExpirationDate" validate:"required"`
    Address string `json:"address" validate:"required"`
    City string `json:"city" validate:"required"`
    State string `json:"state" validate:"required"`
    ZipCode string `json:"zipCode" validate:"required"`
    CarMake string `json:"carMake" validate:"required"`
    CarModel string `json:"carModel" validate:"required"`
    CarYear int `json:"carYear" validate:"required"`
    CarColor string `json:"carColor" validate:"required"`
    CarPlateNumber string `json:"carPlateNumber" validate:"required"`
    CarRegistrationExpirationDate string `json:"carRegistrationExpirationDate" validate:"required"`
    }

