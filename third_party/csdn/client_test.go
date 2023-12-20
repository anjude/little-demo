package csdn

import (
	"github.com/anjude/little-demo/third_party/csdn/schema"
	"testing"
)

func Test_client_GetArticleList(t *testing.T) {
	c := &client{
		userName:  "wx_douxj",
		userToken: "token",
	}
	got, err := c.GetArticleList(schema.GetHotListReq{
		Page:     0,
		PageSize: 10,
		Type:     "",
	})
	if err != nil {
		t.Errorf("GetArticleList() error = %v", err)
		return
	}
	t.Logf("GetArticleList() got = %v", got)
}

func Test_client_SubmitComment(t *testing.T) {
	c := &client{
		userName:  "xx",
		userToken: "xx",
	}
	got, err := c.SubmitComment(schema.CommentReq{
		CommentId: "",
		Content:   "test_me",
		ArticleId: "134756794",
	})
	if err != nil {
		t.Errorf("SubmitComment() error = %v", err)
		return
	}
	t.Logf("SubmitComment() got = %v", got)
}
