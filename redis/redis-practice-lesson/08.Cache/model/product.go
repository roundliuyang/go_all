package model

import "redis-parctice-lesson/global"

type Product struct {
	ID            int64   `gorm:"primary_key;type:int"`
	Name          string  `gorm:"name;type:varchar(64)"`
	Price         float64 `gorm:"price"`
	DiscountPrice float64 `gorm:"discount_price"`
}

func (p Product) Add(name string, price, discountPrice float64) (int64, error) {
	product := Product{
		Name:          name,
		Price:         price,
		DiscountPrice: discountPrice,
	}
	result := global.DB.Save(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return product.ID, nil

}

func (p Product) FindAll() []Product {
	var productList []Product
	global.DB.Find(&productList)
	return productList
}

func (p Product) FindById(id int64) (Product, error) {
	var product Product
	result := global.DB.First(&product, id)
	return product, result.Error
}
