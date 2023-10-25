package wenxin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"official-account-dev/conf"
	"time"
)

func RefreshAccessTokenWithRetry() (string, error) {
	// 简单带重试的实现
	// 1. 重试次数
	retryTimes := 3
	// 2. 重试间隔
	retryInterval := 1
	// 3. 重试策略
	for i := 0; i < retryTimes; i++ {
		if i > 0 {
			// 重试间隔
			time.Sleep(time.Duration(retryInterval) * time.Second)
		}
		token, err := RefreshAccessToken()
		if err != nil {
			fmt.Printf("retry %d times, err: %v\n", i, err)
			continue
		}
		return token, nil
	}
	return "", fmt.Errorf("refresh access token failed")
}

func RefreshAccessToken() (string, error) {
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		conf.Get().APIKey, conf.Get().APISecret)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	result := GetTokenResp{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.AccessToken, nil
}

type GetTokenResp struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session_secret"`
}
