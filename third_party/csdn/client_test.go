package csdn

import (
	"github.com/anjude/little-demo/third_party/csdn/schema"
	"testing"
)

func Test_Demo(t *testing.T) {
	c := NewClient("", "")
	got, err := c.GetArticleList(schema.GetHotListReq{
		Page:     0,
		PageSize: 10,
		Type:     "",
	})
	if err != nil {
		t.Errorf("GetArticleList() error = %v", err)
		return
	}
	for _, article := range got.Data {
		_, err := c.SubmitComment(schema.CommentReq{
			CommentId: "",
			Content:   "请回复有意义的内容",
			ArticleId: article.ProductId,
		})
		if err != nil {
			t.Errorf("SubmitComment() error = %v", err)
			return
		}
	}
}

func Test_client_GetArticleList(t *testing.T) {
	c := NewClient("", "")
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
	c := NewClient("", "")
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
	c := NewClient("", "")
	got, err := c.GetCommentList(schema.CommentListReq{
		ArticleId: "134990421",
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
	c := NewClient("", "")
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

func Test_csdnClient_GetMyArticleCommentList(t *testing.T) {
	c := NewClient("", "")
	got, err := c.GetMyArticleCommentList(schema.ArticleCommentListReq{
		Type: "in",
		Page: 1,
		Size: 20,
	})
	if err != nil {
		t.Errorf("get comment error = %v", err)
		return
	}
	t.Logf("comment got = %v", got)
}

func Test_csdnClient_GetUserArticleList(t *testing.T) {
	c := NewClient("", "")
	got, err := c.GetUserArticleList(schema.GetUserArticleListReq{
		BusinessType: "blog",
		Page:         1,
		Size:         20,
		UserName:     "weixin_52908342",
	})
	if err != nil {
		t.Errorf("get error = %v", err)
		return
	}
	t.Logf("got = %v", got)
}
