package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	// 网页URL
	url := "http://mp.weixin.qq.com/s?__biz=MzI4OTg4NTE3OA==&mid=2247490946&idx=1&sn=dc2c8a20854349d42d61ac97229513ee&chksm=ec290317db5e8a0119e329cd094ee48a52f2b193ac2df3296c32d0003e74581a39f4fa5c79a3#rd"

	// 获取网页内容
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the webpage:", err)
		return
	}
	defer resp.Body.Close()

	// 解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing the HTML document:", err)
		return
	}

	// 查找标题
	var title string
	var author string
	var content string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				for _, a := range n.Attr {
					if a.Key == "id" && a.Val == "activity-name" {
						title = extractText(n)
					}
				}
			case "meta":
				for _, a := range n.Attr {
					if a.Key == "property" && a.Val == "og:title" {
						title = extractText(n)
					}
					if a.Key == "property" && a.Val == "og:author" {
						author = extractText(n)
					}
				}
			case "p":
				content += extractText(n) + "\n"
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	// 输出结果
	fmt.Println("Title:", title)
	fmt.Println("Author:", author)
	fmt.Println("Content:", content)

}

// extractText 从HTML节点中提取文本内容
func extractText(n *html.Node) string {
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			ret += c.Data
		}
	}
	return ret
}
