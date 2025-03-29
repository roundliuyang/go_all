package model

import "time"

type Point struct {
	ID        int64 `gorm:"id;primary_key"`
	OrderId   int64 `gorm:"order_id"`
	Amount    int   `gorm:"amount"`
	Status    int   `gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Point) Add(orderId int64, amount, status int) (Point, error) {
	point := Point{
		OrderId:   orderId,
		Amount:    amount,
		Status:    status,
		CreatedAt: time.Now(),
	}
	result := DB.Save(&point)
	return point, result.Error
}
