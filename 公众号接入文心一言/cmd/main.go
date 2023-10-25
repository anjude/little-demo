package main

import (
	"log"
	"net/http"
	core "official-account-dev/api"
	"official-account-dev/cache"
	"official-account-dev/conf"
)

// @Author 豆小匠Coding
func main() {
	// 初始化缓存
	cache.InitCache()

	// 初始化配置
	conf.InitConf()

	// 提供一个http请求接口
	http.HandleFunc("/official_auto_reply", core.GetAutoReply)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("启动服务器时发生错误:", err)
	}
}
