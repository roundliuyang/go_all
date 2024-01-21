package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Goods struct {
	Id         int
	Title      string
	Price      float64
	Stock      int
	Type       int
	CreateTime time.Time
}

func (v Goods) TableName() string {
	return "goods"
}

func SaveGoods(goods Goods) {
	DB.Create(&goods)
}

// 更新
func UpdateGoods() {
	goods := Goods{}
	DB.Where("id = ?", 1).Take(&goods)
	goods.Price = 1000
	//UPDATE `goods` SET `title`='毛巾',`price`=100.000000,`stock`=100,`type`=0,`create_time  `='2022-11-25 13:03:48' WHERE `id` = 1
	DB.Save(&goods)

	// 更新单个列
	DB.Model(&goods).Update("title", "hello")
	//DB.Model(&Goods{}).Where("id", 1).Update("title", "hello2")

	// 更新多列
	DB.Model(&goods).Updates(Goods{
		Title: "hello3",
		Stock: 200,
	})

	// 更新选定的字段
	DB.Model(&goods).Select("title").Updates(Goods{
		Title: "hello",
		Stock: 200,
	})
}

// 删除
func DeleteGoods() {
	DB.Delete(&Goods{}, 1)
	//DB.Where("id=?", 1).Delete(&Goods{})
}

// 查询
func FindGoods() {
	var goods Goods
	//  SELECT * FROM `goods` WHERE id=1 LIMIT 1
	DB.Where("id=?", 1).Take(&goods)

	//  SELECT * FROM `goods` WHERE id=1 AND `goods`.`id` = 1
	DB.Where("id=?", 1).Find(&goods)

	// SELECT * FROM `goods` WHERE id=1 AND `goods`.`id` = 1 ORDER BY `goods`.`id` LIMIT 1
	DB.Where("id=?", 1).First(&goods)

	// SELECT * FROM `goods` WHERE id=1 AND `goods`.`id` = 1 LIMIT 1
	DB.Where("id=?", 1).Limit(1).Find(&goods)
}

// 分页查询
func FindPageGoods() {
	var goods []Goods
	DB.Order("create_time desc").Limit(10).Offset(10).Find(&goods)
}

// 分组 todo

// 直接执行sql语句
func ExecGoods() {
	//统计每个商品分类下面有多少个商品
	//定一个Result结构体类型，用来保存查询结果
	type Result struct {
		Type  int
		Total int
	}
	var results []Result

	sql := "SELECT type, count(*) as  total FROM `goods` where create_time > ? GROUP BY type HAVING (total > 0)"
	//因为sql语句使用了一个问号(?)作为绑定参数, 所以需要传递一个绑定参数(Raw第二个参数).
	//Raw函数支持绑定多个参数
	DB.Raw(sql, "2022-11-06 00:00:00").Scan(&results)
	fmt.Println(results)
}

// 自动事务
func Transaction() {
	db := DB.Session(&gorm.Session{})
	err := db.Transaction(func(tx *gorm.DB) error {
		// 事务操作
		goods1 := Goods{
			Title: "苹果派",
			Price: 6.5,
			Stock: 200,
			Type:  0,
		}
		if err := tx.Create(&goods1).Error; err != nil {
			return err
		}

		goods2 := Goods{
			Title: "苹果派苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果",
			Price: 6.5,
			Stock: 200,
			Type:  0,
		}
		if err := tx.Create(&goods2).Error; err != nil {
			return err
		}
		return nil
	})
	log.Println("transaction err", err)
}

// 手动事务
func Transaction2() {
	tx := DB.Begin()
	goods1 := Goods{
		Title: "苹果派",
		Price: 6.5,
		Stock: 200,
		Type:  0,
	}
	if err := tx.Create(&goods1).Error; err != nil {
		tx.Rollback()
		return
	}

	goods2 := Goods{
		Title: "苹果派苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果",
		Price: 6.5,
		Stock: 200,
		Type:  0,
	}
	if err := tx.Create(&goods2).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}

// 嵌套事务 todo

// 保存点
func Transaction3() {
	goods1 := Goods{
		Title: "苹果派",
		Price: 6.5,
		Stock: 200,
		Type:  0,
	}
	goods2 := Goods{
		Title: "苹果派苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果苹果",
		Price: 6.5,
		Stock: 200,
		Type:  0,
	}

	tx := DB.Begin()
	tx.Create(&goods1)

	tx.SavePoint("sp1")
	tx.Create(&goods2)
	tx.RollbackTo("sp1") // Rollback user2

	tx.Commit() // Commit user1
}

func (*Goods) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("before create .....")
	return nil
}

func (*Goods) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("after create .....")
	return nil
}

func (*Goods) AfterSave(tx *gorm.DB) (err error) {
	log.Println("after save .....")
	return nil
}

func (*Goods) BeforeSave(tx *gorm.DB) (err error) {
	log.Println("before save .....")
	return nil
}
