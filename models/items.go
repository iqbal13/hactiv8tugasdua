package models

import "time"

type Items struct {
	ItemID      uint      `json:"itemId" gorm:"primaryKey"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"orderId" gorm:"foreignKey:OrderId;References:OrderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
