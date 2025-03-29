package model

import "time"

const (
	DeliveryStatusFree = iota + 1
	DeliveryStatusBusy
)

type Delivery struct {
	ID        int64  `gorm:"id;primary_key"`
	Name      string `gorm:"name"`
	Status    int    `gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d Delivery) FindByStatus(status int) ([]Delivery, error) {
	var deliveryList []Delivery
	result := DB.Model(Delivery{}).Where("status=?", status).Find(&deliveryList)
	return deliveryList, result.Error
}
