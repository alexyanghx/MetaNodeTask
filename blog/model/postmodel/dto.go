package postmodel

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type QueryPostPageDTO struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}

type UpdatePostDTO struct {
	ID      uint   `json:"id" binding:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (dto *UpdatePostDTO) Validate() error {
	valid := validator.New()
	if dto.Title == "" && dto.Content == "" {
		return fmt.Errorf("title 和 content 至少需要填写一个")
	}
	return valid.Struct(dto)
}
