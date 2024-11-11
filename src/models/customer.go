package models

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"Name" binding:"required"`
	BirthDate string `json:"BirthDate" binding:"required"`
	Address   string `json:"Address" binding:"required"`
	Phone     string `json:"Phone" binding:"required"`
}
