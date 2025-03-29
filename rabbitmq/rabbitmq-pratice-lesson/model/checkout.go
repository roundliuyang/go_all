package model

import "time"

type CheckOut struct {
	ID            int64 `gorm:"id;primary_key"`
	OrderId       int64 `gorm:"order_id"`
	TransactionId int64 `gorm:"transaction_id"`
	Amount        int   `gorm:"amount"`
	Status        int   `gorm:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (c CheckOut) Save(orderId, transactionId int64, amount, status int) (CheckOut, error) {
	checkOut := CheckOut{
		OrderId:       orderId,
		TransactionId: transactionId,
		Amount:        amount,
		Status:        status,
		CreatedAt:     time.Now(),
	}
	result := DB.Save(&checkOut)
	return checkOut, result.Error
}
