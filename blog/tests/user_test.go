package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexyanghx/MyBlog/model/usermodel"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestUser(t *testing.T) {

	t.Run("register", func(t *testing.T) {

		resp := httptest.NewRecorder()
		//传入user的数据
		reqBody := gin.H{
			"username": "alex",
			"password": "123456",
			"email":    "alex@163.com",
		}
		body, _ := json.Marshal(reqBody)
		fmt.Println(string(body))
		req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(body))
		r.ServeHTTP(resp, req)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("login", func(t *testing.T) {
		user := usermodel.User{
			Username: "alex",
			Password: "123456",
		}

		body, _ := json.Marshal(&user)

		resp := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

}
