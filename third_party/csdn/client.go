package csdn

import (
	"encoding/json"
	"fmt"
	"github.com/anjude/little-demo/third_party/csdn/schema"
	"io"
	"net/http"
	"strings"
)

type csdnClient struct {
	userName  string
	userToken string
}

func NewClient(u, t string) *csdnClient {
	return &csdnClient{
		userName:  u,
		userToken: t,
	}
}

func (c *csdnClient) GetUserName() string {
	return c.userName
}

func (c *csdnClient) GetUserToken() string {
	return c.userToken
}

func (c *csdnClient) GetArticleList(param schema.GetHotListReq) (*schema.GetHotListResp, error) {
	url := fmt.Sprintf("https://blog.csdn.net/phoenix/web/blog/hot-rank?page=%d&pageSize=%d&type=%s",
		param.Page, param.PageSize, param.Type)

	cli := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := &schema.GetHotListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *csdnClient) SubmitComment(param schema.CommentReq) (*schema.CommentResp, error) {
	url := "https://blog.csdn.net/phoenix/web/v1/comment/submit"

	payload := strings.NewReader(fmt.Sprintf("commentId=%s&content=%s&articleId=%s", param.CommentId, param.Content, param.ArticleId))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Cookie", fmt.Sprintf("UserName=%s; UserToken=%s", c.userName, c.userToken))
	req.Header.Add("Origin", "https://blog.csdn.net")
	req.Header.Add("Referer", "https://blog.csdn.net/weixin_52908342/article/details/134990421")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := &schema.CommentResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *csdnClient) GetCommentList(param schema.CommentListReq) (*schema.CommentListResp, error) {
	url := fmt.Sprintf("https://blog.csdn.net/phoenix/web/v1/comment/list/%s?page=%d&size=%d", param.CommentId, param.Page, param.Size)

	cli := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resp := &schema.CommentListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
