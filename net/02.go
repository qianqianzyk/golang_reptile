package main

//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"regexp"
//	"strings"
//	"sync"
//)
//
//// request 发送HTTP GET请求并返回响应的内容
//func request(url string) (string, error) {
//	// 创建一个http.Client
//	client := &http.Client{}
//
//	// 创建新的HTTP请求
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return "", fmt.Errorf("new request error: %v", err)
//	}
//
//	// 设置Header
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0")
//	req.Header.Add("Cookie", "Hm_lvt_87d678935e4b33455c0390543e7a759d=1719920808; Hm_lpvt_87d678935e4b33455c0390543e7a759d=1719920808")
//
//	// 发出请求
//	resp, err := client.Do(req)
//	if err != nil {
//		return "", fmt.Errorf("http get error: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// 检查响应状态码
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("http status code: %d", resp.StatusCode)
//	}
//
//	// 读取响应体
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return "", fmt.Errorf("read error: %v", err)
//	}
//
//	return string(body), nil
//}
//
//// 解析 HTML 内容，提取链接并发起新的请求
//func parse(html string) {
//	// 移除换行符
//	html = strings.Replace(html, "\n", "", -1)
//
//	// 匹配边栏内容的正则表达式
//	reSidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`)
//	sidebar := reSidebar.FindString(html)
//
//	// 匹配链接的正则表达式
//	reLink := regexp.MustCompile(`href="(.*?)"`)
//	links := reLink.FindAllString(sidebar, -1)
//
//	baseURL := "https://gorm.io/zh_CN/docs/"
//	var wg sync.WaitGroup
//
//	// 遍历所有链接
//	for _, v := range links {
//		s := v[6 : len(v)-1] // 提取链接部分
//		url := baseURL + s
//		fmt.Printf("url: %v\n", url)
//
//		wg.Add(1)
//		go func(url string) {
//			defer wg.Done()
//			body, err := request(url)
//			if err != nil {
//				log.Printf("Failed to fetch URL: %v, error: %v", url, err)
//				return
//			}
//			parse2(body) // 解析新的页面内容
//		}(url)
//	}
//
//	wg.Wait() // 等待所有并发请求完成
//}
//
//// 解析页面内容并提取标题
//func parse2(body string) {
//	// 移除换行符
//	body = strings.Replace(body, "\n", "", -1)
//
//	// 匹配页面内容的正则表达式
//	reContent := regexp.MustCompile(`<div class="article">(.*?)</div>`)
//	content := reContent.FindString(body)
//	if content == "" {
//		fmt.Println("未找到内容")
//		return
//	}
//
//	// 匹配标题的正则表达式
//	reTitle := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
//	title := reTitle.FindString(content)
//	if title == "" {
//		fmt.Println("未找到标题")
//		return
//	}
//
//	// 提取并打印标题
//	title = title[42 : len(title)-5]
//	fmt.Printf("title: %v\n", title)
//}
//
//func main() {
//	url := "https://gorm.io/zh_CN/docs/"
//	s, err := request(url)
//	if err != nil {
//		log.Fatalf("Failed to fetch URL content: %v", err)
//	}
//	parse(s)
//}
