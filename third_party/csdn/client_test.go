package csdn

import (
	"github.com/anjude/little-demo/third_party/csdn/schema"
	"testing"
)

func Test_client_GetArticleList(t *testing.T) {
	c := &csdnClient{
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
	c := &csdnClient{
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

func Test_csdnClient_GetCommentList(t *testing.T) {
	c := &csdnClient{
		userName:  "xxx",
		userToken: "xxx",
	}
	got, err := c.GetCommentList(schema.CommentListReq{
		CommentId: "134990421",
		Page:      1,
		Size:      10,
	})
	if err != nil {
		t.Errorf("GetCommentList() error = %v", err)
		return
	}
	t.Logf("GetCommentList() got = %v", got)
}

func Test_csdnClient_DiggComment(t *testing.T) {
	c := &csdnClient{
		userName:  "xxx",
		userToken: "xxx",
	}
	got, err := c.DiggComment(schema.DiggCommentReq{
		CommentId: "30508799",
		ArticleId: "135147263",
	})
	if err != nil {
		t.Errorf("DiggComment error = %v", err)
		return
	}
	t.Logf("DiggComment got = %v", got)
}
