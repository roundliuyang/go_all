package service

import (
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
	"test/interface_mock/blog"
	mblog "test/interface_mock/test/mocks/blog"
	"testing"
)

func TestListPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlog := mblog.NewMockBlog(ctrl)
	mockBlog.EXPECT().ListPosts().Return([]blog.Post{})

	service := NewService(mockBlog)

	posts, _ := service.ListPosts()       // 获取返回值和 error
	assert.Equal(t, []blog.Post{}, posts) // 然后检查 posts 是否为空数组
}
