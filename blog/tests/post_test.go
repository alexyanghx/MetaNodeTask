package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alexyanghx/MyBlog/middleware"
	"github.com/alexyanghx/MyBlog/model/postmodel"
	"github.com/golang-jwt/jwt"
	"github.com/magiconair/properties/assert"
	"gorm.io/gorm"
)

func TestPost(t *testing.T) {
	t.Run("post create", func(t *testing.T) {
		post := postmodel.Post{
			Title:   "test",
			Content: "test",
		}

		body, _ := json.Marshal(&post)

		resp := httptest.NewRecorder()
		//请求头上需要填充 token

		req := httptest.NewRequest("POST", "/post/create", bytes.NewBuffer(body))
		tokenString, _ := middleware.GenerateToken(&middleware.CustomClaims{
			UserId:   1,
			Roles:    []string{"admin"},
			Username: "alex",
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		})
		req.Header.Set("Authorization", "Bearer "+tokenString)
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("queryPostPage", func(t *testing.T) {
		var queryPostPageDTO = postmodel.QueryPostPageDTO{
			PageNum:  1,
			PageSize: 10,
		}
		body, _ := json.Marshal(queryPostPageDTO)

		req := httptest.NewRequest("POST", "/post/queryPage", bytes.NewBuffer(body))
		setHeaderToken(req, 1, []string{"admin"})
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

	})

	t.Run("postUpdate", func(t *testing.T) {
		var post = postmodel.Post{
			Model: gorm.Model{ID: 1},
			Title: "test456",
		}
		body, _ := json.Marshal(post)

		req := httptest.NewRequest("POST", "/post/update", bytes.NewBuffer(body))
		setHeaderToken(req, 1, []string{"admin"})
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

	})

	t.Run("postDelete", func(t *testing.T) {

		req := httptest.NewRequest("DELETE", "/post/delete/1", nil)
		setHeaderToken(req, 1, []string{"admin"})
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)

	})
}
