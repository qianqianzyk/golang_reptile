package main

//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"time"
//)
//
//// 一个简单爬虫的实现
//// request 发送HTTP GET请求并返回响应的内容
//func request(url string) (string, error) {
//	// 创建一个http.Client
//	client := &http.Client{
//		Timeout: 30 * time.Second, // 设置超时时间
//	}
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
//func main() {
//	url := "https://gorm.io/zh_CN/docs/"
//	content, err := request(url)
//	if err != nil {
//		log.Fatalf("Failed to fetch URL content: %v", err)
//	}
//
//	fmt.Printf("Content: %v\n", content)
//}
