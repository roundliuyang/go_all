package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"redis-parctice-lesson/08.Cache/model"
	"redis-parctice-lesson/global"
	"time"
)

func FindAll() {
	p := model.Product{}
	productList := p.FindAll()
	for _, item := range productList {
		key := fmt.Sprintf("%s%d", global.Product, item.ID)
		v, _ := json.Marshal(item)
		result, err := global.GoRedisClient.Set(context.Background(), key, v, 5*time.Minute).Result()
		if err != nil {
			//TODO处理
		}
		fmt.Println(result)
	}
}

func FindAllByPipeline() {
	p := model.Product{}
	productList := p.FindAll()
	global.GoRedisClient.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		for _, item := range productList {
			key := fmt.Sprintf("%s%d", global.Product, item.ID)
			v, _ := json.Marshal(item)
			result, err := global.GoRedisClient.Set(context.Background(),
				key, v, 5*time.Minute).Result()
			if err != nil {
				//TODO处理
				return err
			}
			fmt.Println(result)
		}
		return nil
	})
}

func AddProduct(name string, price, discountPrice float64) (int64, error) {
	p := &model.Product{}
	r, err := p.Add(name, price, discountPrice)
	return r, err
}

func FindById(productId int64) (model.Product, error) {
	var product model.Product
	if productId < 1 {
		return product, errors.New("参数错误")
	}
	key := fmt.Sprintf("%s%d", global.Product, productId)
	member, err := global.GoRedisClient.Get(context.Background(), key).Result()
	if errors.Is(err, redis.Nil) {
		p := model.Product{}
		product, err := p.FindById(productId)
		if err != nil {
			//TODO
		}
		productJson, err := json.Marshal(product)
		if err != nil {
			return product, err
		}
		_, err = global.GoRedisClient.Set(context.Background(), key, string(productJson), 0).Result()
		if err != nil {
			return product, err
		}
		return product, nil
	}
	err = json.Unmarshal([]byte(member), &product)
	return product, err
}
