package model

import (
	"redis-parctice-lesson/global"
	"time"
)

type VoucherOrder struct {
	ID             int       `json:"id" gorm:"primary_key;type:int"`
	VoucherID      int       `gorm:"voucher_id"`
	AccountId      int       `gorm:"account_id"`
	Status         int       `json:"status" gorm:"status"`         //-1 已取消 0 未支付 1 已支付 2 已消费 3 已过期
	Payment        int       `json:"payment" gorm:"payment"`       //0=微信 1=支付宝
	OrderType      int       `json:"order_type" gorm:"order_type"` //订单类型0=正常订单，1 抢购订单
	CreateDateTime time.Time `gorm:"createDate;default:null"`
	UpdateDateTime time.Time `gorm:"updateDate;default:null"`
	IsValid        int       `gorm:"is_valid"`
}

func (vo *VoucherOrder) Add(voucherId, accountId, status, payment, orderType int) (int, error) {
	voucherOrder := VoucherOrder{
		VoucherID:      voucherId,
		AccountId:      accountId,
		Status:         status,
		Payment:        payment,
		OrderType:      orderType,
		IsValid:        1,
		CreateDateTime: time.Now(),
	}

	result := global.DB.Save(&voucherOrder)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return voucherOrder.ID, nil
	}
}

func (vo VoucherOrder) GetVoucherOrderByCondition(accountId, voucherId int) (VoucherOrder, error) {
	var v VoucherOrder
	result := global.DB.Model(VoucherOrder{}).Where("account_id=?", accountId).Where("voucher_id=?", voucherId).Find(&v)
	if result.Error != nil {
		return v, result.Error
	}
	return v, nil
}
