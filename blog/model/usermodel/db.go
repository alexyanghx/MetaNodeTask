package usermodel

import (
	"github.com/alexyanghx/MyBlog/model"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
}

func (user *User) CreateUser() error {
	return model.DB.Create(user).Error
}

func (user *User) UpdateUser() error {
	return model.DB.Updates(user).Error
}

func QueryUserByName(username string) (*User, error) {
	var user User
	if err := model.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}
