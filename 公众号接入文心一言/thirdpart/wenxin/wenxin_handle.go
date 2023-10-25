package wenxin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"official-account-dev/cache"
	core "official-account-dev/constant"
	"strings"
	"time"
)

type UserType string

const (
	RoleUser      UserType = "user"
	RoleAssistant UserType = "assistant"
)

type ReplyReq struct {
	Messages []Message `json:"messages"`
	System   string    `json:"system"` // 系统预设
}

type ReplyResp struct {
	Id               string `json:"id"`
	Object           string `json:"object"`
	Created          int    `json:"created"`
	Result           string `json:"result"`
	IsTruncated      bool   `json:"is_truncated"`
	NeedClearHistory bool   `json:"need_clear_history"`
	Usage            struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Message struct {
	Role    UserType `json:"role"`
	Content string   `json:"content"`
}

// GetWenXinReply 请求文心一言，多轮对话，非流式回复
func GetWenXinReply(msgs []Message) (string, error) {
	wenxinToken, ok := cache.Get(core.GetCacheKey(core.CacheWenXinToken))
	if !ok {
		var err error
		wenxinToken, err = RefreshAccessTokenWithRetry()
		if err != nil {
			return "", err
		}
		// 缓存29天(access token有效期为30天
		cache.Set(core.GetCacheKey(core.CacheWenXinToken), wenxinToken, 29*time.Hour*24)
	}
	url := fmt.Sprintf("https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=%s", wenxinToken)
	method := "POST"

	marshal, err := json.Marshal(ReplyReq{msgs, "你是一名塔罗牌占卜师，请咨询我有什么疑问！"})
	if err != nil {
		return "", err
	}
	payload := strings.NewReader(string(marshal))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	result := ReplyResp{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.Result + " 长时间未回复，输入”喵“即可获取答案哦！", nil
}
