package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alexyanghx/MyBlog/middleware"
	"github.com/alexyanghx/MyBlog/model/commentmodel"
	"github.com/alexyanghx/MyBlog/model/postmodel"
	"github.com/alexyanghx/MyBlog/model/usermodel"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Result struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Register(c *gin.Context) {
	var user usermodel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &Result{Msg: "参数解析异常"})
		return
	}

	queryUser, err := usermodel.QueryUserByName(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Result{Msg: "查询异常"})
		return
	}

	if queryUser != nil {
		c.JSON(http.StatusBadRequest, &Result{Msg: "用户名已存在"})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Result{Msg: "密码加密异常"})
		return
	}
	user.Password = string(password)
	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusInternalServerError, &Result{Msg: "创建用户异常"})
		return
	}
	c.JSON(http.StatusOK, &Result{Msg: "注册用户成功", Data: user})

}

func Login(c *gin.Context) {
	var user usermodel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, &Result{Msg: "参数解析异常"})
		return
	}

	var queryUser *usermodel.User
	var err error
	if queryUser, err = usermodel.QueryUserByName(user.Username); err != nil {
		c.JSON(http.StatusInternalServerError, &Result{Msg: "查询用户异常"})
		return
	}

	if queryUser == nil {
		c.JSON(http.StatusBadRequest, &Result{Msg: "用户不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(queryUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, &Result{Msg: "密码错误"})
		return
	}

	token, err := middleware.GenerateToken(&middleware.CustomClaims{
		Roles:    []string{"admin"},
		UserId:   queryUser.ID,
		Username: queryUser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	fmt.Printf("token=%s\n", token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &Result{Msg: "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, &Result{Msg: "登录成功", Data: &map[string]interface{}{
		"token":    token,
		"username": queryUser.Username,
		"userId":   queryUser.ID,
	}})

}

func CreatePost(c *gin.Context) {
	var post postmodel.Post
	if err := c.ShouldBindBodyWithJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "用户不存在"})
		return
	}
	post.UserID = userId.(uint)

	if err := post.CreatePost(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "创建文章失败"})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, &Result{Msg: "创建文章成功"})
}

func QueryPostPage(c *gin.Context) {
	var queryPage postmodel.QueryPostPageDTO

	if err := c.ShouldBindBodyWithJSON(&queryPage); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}

	posts, err := postmodel.QueryPostPage(&queryPage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &Result{Msg: "查询文章失败"})
		return
	}

	c.JSON(http.StatusOK, &Result{Data: posts, Msg: "查询成功"})

}

func UpdatePost(c *gin.Context) {
	var updatePostDTO postmodel.UpdatePostDTO
	if err := c.ShouldBindBodyWithJSON(&updatePostDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}

	if err := updatePostDTO.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: err.Error()})
		return
	}

	userId, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "用户不存在"})
		return
	}

	userID := userId.(uint)
	queryPost, err := postmodel.QueryPostById(updatePostDTO.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &Result{Msg: "文章不存在"})
		return
	}

	if queryPost.UserID != userID {
		c.AbortWithStatusJSON(http.StatusForbidden, &Result{Msg: "无权限"})
		return
	}

	var post postmodel.Post = postmodel.Post{
		Model:   gorm.Model{ID: updatePostDTO.ID},
		UserID:  userID,
		Title:   updatePostDTO.Title,
		Content: updatePostDTO.Content,
	}

	if err := post.UpdatePost(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &Result{Msg: "更新文章失败"})
		return
	}

	c.JSON(http.StatusOK, &Result{Msg: "更新文章成功"})

}

func DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}

	queryPost, err := postmodel.QueryPostById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &Result{Msg: "文章不存在"})
		return
	}

	userId, _ := c.Get("userId")

	if queryPost.UserID != userId.(uint) {
		c.AbortWithStatusJSON(http.StatusForbidden, &Result{Msg: "无权限"})
		return
	}

	if err := queryPost.DeletePost(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &Result{Msg: "删除文章失败"})
		return
	}

	c.JSON(http.StatusOK, &Result{Msg: "删除文章成功"})

}

func CreateComment(c *gin.Context) {
	var createCommentDTO commentmodel.CreateCommentDTO
	if err := c.ShouldBindJSON(&createCommentDTO); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}

	userId, _ := c.Get("userId")
	comment := commentmodel.Comment{
		Content: createCommentDTO.Content,
		PostID:  createCommentDTO.PostId,
		UserID:  userId.(uint),
	}

	if err := comment.CreateComment(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &Result{Msg: "创建评论失败"})
		return

	}

	c.JSON(http.StatusOK, &Result{Msg: "创建评论成功"})
}

func QueryCommentPage(c *gin.Context) {
	var q commentmodel.QueryCommentPageDTO
	if err := c.ShouldBindJSON(&q); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &Result{Msg: "参数错误"})
		return
	}

	comments, err := commentmodel.QueryCommentPage(&q)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &Result{Msg: "查询评论失败"})
		return
	}

	c.JSON(http.StatusOK, &Result{Data: comments})
}
