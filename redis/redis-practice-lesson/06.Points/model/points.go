package model

import (
	"redis-parctice-lesson/global"
	"time"
)

type AccountPoints struct {
	ID         int64     `gorm:"primary_key;type:int"`
	AccountId  int       `gorm:"account_id"`
	Points     int       `gorm:"points"`
	Kind       int       `gorm:"kind"` // 比如，签到得积分，评论得积分
	IsValid    int       `gorm:"is_valid"`
	CreateDate time.Time `gorm:"create_date;default:null"`
	UpdateDate time.Time `gorm:"update_date;default:null"`
}

func (ac AccountPoints) Add(accountId, points, kind int) (int64, error) {
	ap := AccountPoints{
		AccountId:  accountId,
		Points:     points,
		Kind:       kind,
		IsValid:    1,
		CreateDate: time.Now(),
		UpdateDate: time.Time{},
	}
	result := global.DB.Save(&ap)
	return ap.ID, result.Error
}
