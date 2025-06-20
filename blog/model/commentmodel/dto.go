package commentmodel

type QueryCommentPageDTO struct {
	PageNum  int  `json:"pageNum"`
	PageSize int  `json:"pageSize"`
	PostId   uint `json:"postId"`
}

type CreateCommentDTO struct {
	PostId  uint   `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}
