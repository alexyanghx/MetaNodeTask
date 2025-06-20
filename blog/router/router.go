package router

import (
	"github.com/alexyanghx/MyBlog/controller"
	"github.com/alexyanghx/MyBlog/middleware"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GlobalException())
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	postGroup := r.Group("/post", middleware.JwtParse(), middleware.CheckPrevilege("admin"))
	{
		postGroup.POST("/create", controller.CreatePost)
		postGroup.POST("/queryPage", controller.QueryPostPage)
		postGroup.POST("/update", controller.UpdatePost)
		postGroup.DELETE("/delete/:id", controller.DeletePost)
	}

	commentGroup := r.Group("/comment", middleware.JwtParse(), middleware.CheckPrevilege("admin"))
	{
		commentGroup.POST("/create", controller.CreateComment)
		commentGroup.POST("/queryPage", controller.QueryCommentPage)
	}

	return r
}
