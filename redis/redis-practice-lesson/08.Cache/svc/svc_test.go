package svc

import (
	"fmt"
	"redis-parctice-lesson/08.Cache/model"
	"redis-parctice-lesson/global"
	"testing"
	"time"
)

func init() {
	global.InitDB()
	model.InitDB()
	global.InitGoRedisClient()
}

func TestAddProduct(t *testing.T) {
	for i := 0; i <= 200000; i++ {
		name := fmt.Sprintf("Product-%d", i)
		AddProduct(name, float64(i), float64(i))
	}
}

func TestFindById(t *testing.T) {
	id := 100
	p, err := FindById(int64(id))
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func TestFindAll(t *testing.T) {
	start := time.Now()
	FindAll()
	timeSpan := time.Now().Sub(start)
	fmt.Println(timeSpan)
	//20万条数据
	//3m16.755020458s
}

func TestFindAllByPipeline(t *testing.T) {
	start := time.Now()
	FindAllByPipeline()
	timeSpan := time.Now().Sub(start)
	fmt.Println(timeSpan)
	//20万条数据
	//3.701215417s(3.70s)
}
