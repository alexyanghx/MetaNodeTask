package postmodel

import (
	"github.com/alexyanghx/MyBlog/model"
	"github.com/alexyanghx/MyBlog/model/usermodel"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" binding:"required" json:"title"`
	Content string `gorm:"not null" binding:"required" json:"content"`
	UserID  uint
	User    usermodel.User
}

func (post *Post) CreatePost() error {
	return model.DB.Create(post).Error
}

func (post *Post) UpdatePost() error {
	return model.DB.Updates(post).Error
}

func (post *Post) DeletePost() error {
	return model.DB.Delete(post).Error
}

func QueryPostPage(q *QueryPostPageDTO) ([]Post, error) {
	var posts []Post
	err := model.DB.Preload("User").Limit(q.PageSize).Offset((q.PageNum - 1) * q.PageSize).Find(&posts).Error
	return posts, err
}

func QueryPostById(id uint) (Post, error) {
	var post Post
	err := model.DB.Where("id = ?", id).First(&post).Error
	return post, err
}
