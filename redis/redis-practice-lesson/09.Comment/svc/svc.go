package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"redis-parctice-lesson/09.Comment/model"
	"redis-parctice-lesson/global"
)

func AddComments(productId, like int, content string) (int, error) {
	c := model.Comment{}
	comment, err := c.Add(productId, like, content)
	if err != nil {
		return 0, err
	}
	if comment.ID < 1 {
		return 0, errors.New("数据库添加失败")
	}
	commentJson, err := json.Marshal(comment)
	if err != nil {
		return 0, errors.New("序列化失败")
	}
	key := fmt.Sprintf("%s%d", global.ProductComment, productId)
	result, err := global.GoRedisClient.LPush(context.Background(), key, commentJson).Result()
	if err != nil {
		return 0, err
	}
	fmt.Println(result)
	return comment.ID, nil
}

func FindLatestComments(productId int, start, end int64) ([]model.Comment, error) {
	var comments []model.Comment
	key := fmt.Sprintf("%s%d", global.ProductComment, productId)
	result, err := global.GoRedisClient.LRange(context.Background(),
		key, start, end).Result()
	//TODO 如果redis没有，可以再从数据库里拿
	if err != nil {
		return nil, err
	}
	for _, item := range result {
		var c model.Comment
		json.Unmarshal([]byte(item), &c)
		comments = append(comments, c)
	}
	return comments, nil
}
