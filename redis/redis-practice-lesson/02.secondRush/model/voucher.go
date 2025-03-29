package model

import (
	"redis-parctice-lesson/global"
	"time"
)

//秒杀活动（后台添加的）

type Voucher struct {
	ID         int       `json:"id" gorm:"primary_key;type:int"`
	Amount     int       `json:"amount" gorm:"amount"`
	StartTime  time.Time `json:"startTime" gorm:"start_time"`
	EndTime    time.Time `json:"endTime" gorm:"end_time"`
	CreateTime time.Time `gorm:"create_time;default:null"`
	UpdateTime time.Time `gorm:"update_time;default:null"`
	IsValid    int       `gorm:"is_valid"`
}

func (r Voucher) GetById(id int) (Voucher, error) {
	var v = Voucher{
		ID: id,
	}
	result := global.DB.First(&v)
	if result.Error != nil {
		return v, result.Error
	} else {
		return v, nil
	}
}

func (r Voucher) Add(amount int, startTime, endTime time.Time) (int, error) {
	voucher := Voucher{
		Amount:     amount,
		StartTime:  startTime,
		EndTime:    endTime,
		IsValid:    1,
		CreateTime: time.Now(),
	}
	result := global.DB.Save(&voucher)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return voucher.ID, nil
	}
}

func (r Voucher) DecreaseStock(id int) (int, error) {
	result, err := r.GetById(id)
	if err != nil {
		//TODO 记录日志
		return 0, err
	}
	global.DB.Model(Voucher{}).Where("id=?", id).Update("amount", int32(result.Amount-1))
	return id, nil
}
