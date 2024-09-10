package models

import "time"

type Order struct {
	OrderId      uint      `json:"orderId" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Items        []Items   `json:"items" gorm:"foreignKey:OrderId;References:OrderId"`
}
