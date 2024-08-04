package main

//func main() {
//	// 目标页面的 URL
//	url := "https://gorm.io/zh_CN/docs/"
//
//	// 发起 HTTP GET 请求获取页面内容
//	resp, err := http.Get(url)
//	if err != nil {
//		log.Fatalf("Failed to fetch URL: %v", err)
//	}
//	defer resp.Body.Close()
//
//	// 使用 goquery 解析 HTML 页面
//	doc, err := goquery.NewDocumentFromReader(resp.Body)
//	if err != nil {
//		log.Fatalf("Failed to parse document: %v", err)
//	}
//
//	// 基础 URL，用于构建完整的详细页面 URL
//	baseURL := "https://gorm.io/zh_CN/docs/"
//
//	// 使用 WaitGroup 来等待所有 goroutine 完成
//	var wg sync.WaitGroup
//
//	// 互斥锁，保护共享资源
//	var mu sync.Mutex
//
//	// 遍历侧边栏链接
//	doc.Find(".sidebar-link").Each(func(i int, s *goquery.Selection) {
//		// 增加 WaitGroup 计数
//		wg.Add(1)
//
//		go func(s *goquery.Selection) {
//			// 函数退出时减少 WaitGroup 计数
//			defer wg.Done()
//
//			// 获取链接地址
//			link, exists := s.Attr("href")
//			if !exists {
//				log.Printf("Href attribute not found for element %d", i)
//				return
//			}
//			detailURL := baseURL + link
//
//			// 发起 HTTP GET 请求获取详细页面内容
//			respDetail, err := http.Get(detailURL)
//			if err != nil {
//				log.Printf("Failed to fetch detail URL %s: %v", detailURL, err)
//				return
//			}
//			defer respDetail.Body.Close()
//
//			// 使用 goquery 解析详细页面内容
//			detailDoc, err := goquery.NewDocumentFromReader(respDetail.Body)
//			if err != nil {
//				log.Printf("Failed to parse detail document %s: %v", detailURL, err)
//				return
//			}
//
//			// 获取文章标题和内容
//			title := detailDoc.Find(".article-title").Text()
//			content, err := detailDoc.Find(".article").Html()
//			if err != nil {
//				log.Printf("Failed to retrieve content for %s: %v", detailURL, err)
//				return
//			}
//
//			// 使用互斥锁保护共享资源
//			mu.Lock()
//			defer mu.Unlock()
//
//			// 打印详细页面 URL、标题和内容（示例中为打印，实际应用中可以根据需求处理）
//			fmt.Printf("Detail URL: %v\n", detailURL)
//			fmt.Printf("Title: %v\n", title)
//			fmt.Printf("Content: %v\n", content)
//		}(s)
//	})
//
//	// 等待所有 goroutine 完成
//	wg.Wait()
//}
