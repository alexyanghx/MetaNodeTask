package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexyanghx/MyBlog/model/commentmodel"
	"github.com/magiconair/properties/assert"
)

func TestComment(t *testing.T) {

	for i := 0; i < 1; i++ {
		t.Run("createComment", func(t *testing.T) {

			var comment = commentmodel.CreateCommentDTO{
				Content: fmt.Sprintf("test_%d", i+1),
				PostId:  2,
			}

			body, _ := json.Marshal(&comment)

			fmt.Println(string(body))

			resp := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/comment/create", bytes.NewBuffer(body))
			setHeaderToken(req, 1, []string{"admin"})
			r.ServeHTTP(resp, req)
			assert.Equal(t, resp.Code, http.StatusOK)
		})
	}

	// t.Run("queryCommentPage", func(t *testing.T) {
	// 	var queryCommentPageDTO = commentmodel.QueryCommentPageDTO{
	// 		PageSize: 10,
	// 		PageNum:  1,
	// 		PostId:   2,
	// 	}

	// 	body, _ := json.Marshal(&queryCommentPageDTO)

	// 	resp := httptest.NewRecorder()
	// 	req := httptest.NewRequest("POST", "/comment/queryPage", bytes.NewBuffer(body))
	// 	setHeaderToken(req, 1, []string{"admin"})
	// 	r.ServeHTTP(resp, req)
	// 	assert.Equal(t, resp.Code, http.StatusOK)
	// })
}
