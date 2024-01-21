package dao

import "time"

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
