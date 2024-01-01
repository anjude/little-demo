package csdn

import (
	"encoding/json"
	"fmt"
	"github.com/anjude/bcore/pkg/httputil"
	"github.com/anjude/little-demo/third_party/csdn/schema"
	"io"
	"net/http"
	"strings"
)

var defaultAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

type Client struct {
	httpClient *httputil.HTTPClient
	userName   string
	userToken  string
}

func NewClient(u, t string) *Client {
	return &Client{
		httpClient: httputil.NewHTTPClient("https://blog.csdn.net", 0, defaultAgent),
		userName:   u,
		userToken:  t,
	}
}

func (c *Client) GetUserName() string {
	return c.userName
}

func (c *Client) GetUserToken() string {
	return c.userToken
}

func (c *Client) GetArticleList(param schema.GetHotListReq) (*schema.GetHotListResp, error) {
	resp := &schema.GetHotListResp{}
	err := c.httpClient.Request(http.MethodGet, fmt.Sprintf("/phoenix/web/blog/hot-rank?page=%d&pageSize=%d&type=%s",
		param.Page, param.PageSize, param.Type), nil, nil, resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("resp.Code != 200, resp:%v", resp)
	}
	return resp, nil
}

func (c *Client) SubmitComment(param schema.CommentReq) (*schema.CommentResp, error) {
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
	if resp.Code != 200 {
		return nil, fmt.Errorf("resp.Code != 200, resp:%v", resp)
	}
	return resp, nil
}

func (c *Client) GetCommentList(param schema.CommentListReq) (*schema.CommentListResp, error) {
	url := fmt.Sprintf("https://blog.csdn.net/phoenix/web/v1/comment/list/%s?page=%d&size=%d", param.ArticleId, param.Page, param.Size)

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

func (c *Client) DiggComment(param schema.DiggCommentReq) (*schema.DiggCommentResp, error) {
	url := "https://blog.csdn.net/phoenix/web/v1/comment/digg"

	payload := strings.NewReader(fmt.Sprintf("articleId=%s&commentId=%s", param.ArticleId, param.CommentId))

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
	resp := &schema.DiggCommentResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetMyArticleCommentList(param schema.ArticleCommentListReq) (*schema.ArticleCommentListResp, error) {
	url := fmt.Sprintf("https://bizapi.csdn.net/blog/phoenix/console/v1/comment/list?type=%s&page=%d&size=%d", param.Type, param.Page, param.Size)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("Cookie", fmt.Sprintf("UserName=%s; UserToken=%s", c.userName, c.userToken))
	req.Header.Add("Origin", "https://blog.csdn.net")
	req.Header.Add("Referer", "https://blog.csdn.net/weixin_52908342/article/details/134990421")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Add("x-ca-key", "203803574")
	req.Header.Add("x-ca-nonce", "09a249e4-4760-4206-8be0-fe6b230f8b8e")
	req.Header.Add("x-ca-signature", "abtDlfxMzbl80JwBME0uyvpmbR4aHAPLprwq/JPG5vM=")
	req.Header.Add("x-ca-signature-headers", "x-ca-key,x-ca-nonce")

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
	resp := &schema.ArticleCommentListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetUserArticleList(param schema.GetUserArticleListReq) (*schema.GetUserArticleListResp, error) {
	url := fmt.Sprintf("https://blog.csdn.net/community/home-api/v1/get-business-list?page=%d&size=%d&businessType=%s&username=%s", param.Page, param.Size, param.BusinessType, param.UserName)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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
	resp := &schema.GetUserArticleListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
