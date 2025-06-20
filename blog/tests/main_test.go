package tests

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexyanghx/MyBlog/middleware"
	"github.com/alexyanghx/MyBlog/model"
	"github.com/alexyanghx/MyBlog/model/commentmodel"
	"github.com/alexyanghx/MyBlog/model/postmodel"
	"github.com/alexyanghx/MyBlog/model/usermodel"
	"github.com/alexyanghx/MyBlog/router"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var r *gin.Engine

func setHeaderToken(req *http.Request, userId uint, roles []string) {
	tokenString, _ := middleware.GenerateToken(&middleware.CustomClaims{
		UserId:   userId,
		Roles:    roles,
		Username: "alex",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	req.Header.Set("Authorization", "Bearer "+tokenString)
}

func BeforeTest() {
	r = router.Setup()
	model.SetUpDatabase("../data/blog.db")
	model.DB.AutoMigrate(&usermodel.User{}, &postmodel.Post{}, &commentmodel.Comment{})
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	BeforeTest()
	os.Exit(m.Run())
}
