# 关注公众号进群交流
<img src="doc/OfficialAccounts.png" alt="Example Image" height="150px">

## CSDN 接口使用
### 1. 获取CSDN的user name和user token
打开csdn，打开控制台 - Application - Cookies，找到domain为blog.csdn.net的cookie，复制user_name和user_token的值
![CSDN Cookie](doc/img.png)
### 2. 调用接口
```go
func Test_Demo(t *testing.T) {
	c := &Client{
		userName:  "your user name",
		userToken: "your user token",
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
```
sdk调用
```azure
package test

import (
	"github.com/anjude/bean-sdk-go/beansdk"
	"github.com/anjude/bean-sdk-go/services/csdn_service"
	"testing"
)

func TestSDK(t *testing.T) {
	// visitor token有限
	client := beansdk.NewClient("visitor", "")
	csdnService := csdn_service.NewCsdnService(client, "user name", "user token")
	resp, err := csdnService.HotArticleComment()
	if err != nil {
		t.Errorf("got err: %v", err)
		return
	}
	t.Logf("success: %v", resp)
}
```
### 3. 效果
打开csdn https://mp.csdn.net/mp_blog/manage/comment - 我发表的评论，即可看到已发布的评论
其他接口使用请参考third_party/csdn/client_test.go

注：本代码仅用于学习目的，请不要滥用，否则后果自负
