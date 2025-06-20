package commentmodel

import (
	"github.com/alexyanghx/MyBlog/model"
	"github.com/alexyanghx/MyBlog/model/postmodel"
	"github.com/alexyanghx/MyBlog/model/usermodel"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content" binding:"required"`
	UserID  uint
	User    usermodel.User
	PostID  uint `json:"postId" binding:"required"`
	Post    postmodel.Post
}

func (comment *Comment) CreateComment() error {
	return model.DB.Create(comment).Error
}

func (comment *Comment) UpdateComment() error {
	return model.DB.Updates(comment).Error
}

func (comment *Comment) DeleteComment() error {
	return model.DB.Delete(comment).Error
}

func QueryCommentPage(q *QueryCommentPageDTO) ([]Comment, error) {
	var comments []Comment
	query := model.DB.Preload("User").Preload("Post").Limit(q.PageSize).Offset((q.PageNum - 1) * q.PageSize)
	if q.PostId != 0 {
		query = query.Where("post_id = ?", q.PostId)
	}
	err := query.Find(&comments).Error
	return comments, err
}
