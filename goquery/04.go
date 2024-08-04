package main

//func main() {
//	// 要抓取的网站URL
//	url := "https://gorm.io/zh_CN/docs/"
//
//	// 获取URL
//	resp, err := http.Get(url)
//	if err != nil {
//		log.Fatalf("获取URL失败: %v", err)
//	}
//	defer resp.Body.Close() // 确保使用完后关闭响应体
//
//	// 检查HTTP错误
//	if resp.StatusCode != 200 {
//		log.Fatalf("获取URL失败，状态码: %d", resp.StatusCode)
//	}
//
//	// 加载HTML文档
//	dom, err := goquery.NewDocumentFromReader(resp.Body)
//	if err != nil {
//		log.Fatalf("解析HTML失败: %v", err)
//	}
//
//	// 查找每个侧边栏链接并打印它的href和文本
//	dom.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
//		href, exists := s.Attr("href")
//		if exists {
//			text := s.Text()
//			fmt.Println(i, href, text)
//		}
//	})
//}
