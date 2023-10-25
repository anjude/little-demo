package conf

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var c *config

type config struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func InitConf() {
	// 获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get work directory path:", err)
	}

	// 读取根目录下的文件
	filePath := filepath.Join(wd, "config.json")

	// 读取JSON文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 解析JSON数据到config结构体
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
	}
}

func Get() *config {
	return c
}
