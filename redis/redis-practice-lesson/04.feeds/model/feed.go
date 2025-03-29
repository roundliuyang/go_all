package model

import (
	"gorm.io/gorm"
	"redis-parctice-lesson/global"
	"time"
)

type Feed struct {
	ID           int       `gorm:"primary_key;type:int"`
	AccountId    int       `gorm:"account_id"`
	Content      string    `gorm:"content;type:varchar(128)"`
	AgreeTotal   int       `gorm:"agree_total"`
	CommentTotal int       `gorm:"comment_total"`
	CreateDate   time.Time `gorm:"create_date;default:null"`
	UpdateDate   time.Time `gorm:"update_date;default:null"`
	Deleted      gorm.DeletedAt
	IsValid      int `gorm:"is_valid"`
}

func (f Feed) Add(accountId, agreeTotal, commentTotal, isValid int, content string) (int, error) {
	feed := Feed{
		AccountId:    accountId,
		AgreeTotal:   agreeTotal,
		CommentTotal: commentTotal,
		IsValid:      isValid,
		Content:      content,
		CreateDate:   time.Now(),
	}
	result := global.DB.Save(&feed)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return feed.ID, nil
	}
}

func (f Feed) FindById(id int) Feed {
	var feed Feed
	global.DB.Model(Feed{}).Where("id=? and is_valid=1", id).Find(&feed)
	return feed
}

func (f Feed) DeletedById(id int) {
	feed := Feed{ID: id}
	global.DB.Delete(&feed)
}

func (f Feed) FindByAccountId(accountId int) []Feed {
	var feedList []Feed
	global.DB.Model(Feed{}).Where("account_id=?", accountId).Find(&feedList)
	return feedList
}

func (f Feed) FindFeedsByIds(feedIds []string) []Feed {
	var feedList []Feed
	global.DB.Where("id in ?", feedIds).Find(&feedList)
	return feedList
}
