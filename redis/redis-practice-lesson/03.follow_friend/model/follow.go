package model

import (
	"redis-parctice-lesson/global"
	"time"
)

type Follow struct {
	ID              int       `gorm:"primary_key"`
	AccountId       int       `gorm:"account_id"`
	FollowAccountId int       `gorm:"follow_account_id"`
	IsValid         int       `gorm:"is_valid"`
	CreateDate      time.Time `gorm:"create_date;default:null"`
	UpdateDate      time.Time `gorm:"update_date;default:null"`
}

func (f Follow) SelectFollow(accountId, followAccountId int) Follow {
	var follow Follow
	global.DB.Model(Follow{}).Where("account_id=?", accountId).
		Where("follow_account_id=?", followAccountId).Find(&follow)
	return follow
}

func (f Follow) Add(accountId, followAccountId int) (int, error) {
	follow := Follow{
		AccountId:       accountId,
		FollowAccountId: followAccountId,
		IsValid:         1,
		CreateDate:      time.Now(),
	}
	result := global.DB.Save(&follow)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return f.ID, nil
	}
}

func (f Follow) Update(id, isValid int) (int, error) {
	now := time.Now()
	r := global.DB.Model(Follow{}).Where("id=?", id).Updates(Follow{IsValid: isValid, UpdateDate: now})
	if r.Error != nil {
		return 0, r.Error
	} else {
		return id, nil
	}
}
