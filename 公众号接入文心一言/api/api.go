package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"official-account-dev/cache"
	"official-account-dev/constant"
	"official-account-dev/thirdpart/wenxin"
	"strings"
	"time"
)

// GetAutoReply 获取自动回复
func GetAutoReply(w http.ResponseWriter, r *http.Request) {
	req := &AutoReplyReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		_, _ = w.Write([]byte("decode body error"))
	}

	resp, err := getReply(req)
	if err != nil {
		_, _ = w.Write([]byte("get reply error"))
	}

	// 返回JSON响应给前端
	jsonData, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

func getReply(req *AutoReplyReq) (*AutoReplyRsp, error) {
	resp := &AutoReplyRsp{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "喵喵喵?",
	}
	if req.Content == "" {
		return nil, nil
	}
	var err error
	reply := ""
	var messages []wenxin.Message
	messageStr, ok := cache.Get(fmt.Sprintf(constant.GetCacheKey(constant.CacheWenXinMsg), req.FromUserName))
	if ok {
		//解析messageStr为messages
		err = json.Unmarshal(messageStr.([]byte), &messages)
		if err != nil {
			return nil, err
		}

		// 如果req.Content开头是喵，直接返回message的最后一个元素
		if strings.HasPrefix(req.Content, "喵") && len(messages) > 0 {
			lastMsg := messages[len(messages)-1]
			if lastMsg.Role == wenxin.RoleUser {
				resp.Content = "还没出来了啦！"
				return resp, nil
			}
			resp.Content = lastMsg.Content
			// 记录用户使用次数到缓存

			return resp, nil
		}
		// 只保留最近的10条记录
		if len(messages) > 10 {
			messages = messages[len(messages)-10:]
		}
	}
	messages = append(messages, wenxin.Message{
		Role:    wenxin.RoleUser,
		Content: req.Content,
	})
	setUserMsgCache(req.FromUserName, messages)
	reply, err = wenxin.GetWenXinReply(messages)
	if err != nil {
		log.Printf("something went wrong. resp: %+v, err: %v", resp, err)
		reply = "获取回答出错啦~"
	} else {
		messages = append(messages, wenxin.Message{
			Role:    wenxin.RoleAssistant,
			Content: reply,
		})
		setUserMsgCache(req.FromUserName, messages)
	}
	resp.Content = reply
	return resp, nil
}

func setUserMsgCache(fromUserName string, messages []wenxin.Message) {
	marshal, err := json.Marshal(messages)
	if err == nil {
		cache.Set(fmt.Sprintf(constant.GetCacheKey(constant.CacheWenXinMsg), fromUserName), marshal, 12*time.Hour)
	}
}
