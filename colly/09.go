package main

//func main() {
//	// 创建一个新的 Colly 收集器
//	c := colly.NewCollector(
//		colly.MaxDepth(2), // 限制抓取深度为2，防止过深的抓取导致资源浪费
//	)
//
//	// 设置最大并发请求数量为2，每个请求之间的延迟为1秒
//	c.Limit(&colly.LimitRule{
//		DomainGlob:  "*",             // 规则适用于所有域名
//		Parallelism: 2,               // 最大并发请求数
//		Delay:       1 * time.Second, // 每个请求之间的延迟
//	})
//
//	// 在每个 .sidebar-link 元素上执行的回调函数
//	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
//		href := e.Attr("href") // 获取链接的 href 属性
//		// 排除 index.html 以避免重复访问首页
//		if href != "index.html" {
//			// 转换为绝对 URL 并访问该链接
//			c.Visit(e.Request.AbsoluteURL(href))
//		}
//	})
//
//	// 在每个 .article-title 元素上执行的回调函数
//	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
//		title := h.Text              // 获取文章标题文本
//		fmt.Println("Title:", title) // 打印文章标题
//	})
//
//	// 在每个 .article 元素上执行的回调函数
//	c.OnHTML(".article", func(h *colly.HTMLElement) {
//		content, err := h.DOM.Html() // 获取文章内容的 HTML
//		if err != nil {
//			log.Println("Error getting article content:", err)
//		}
//		fmt.Println("Content:", content) // 打印文章内容
//	})
//
//	// 在每个请求发出之前执行的回调函数
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting", r.URL.String()) // 打印即将访问的 URL
//	})
//
//	// 在每个请求发生错误时执行的回调函数
//	c.OnError(func(r *colly.Response, err error) {
//		log.Printf("Request to URL %v failed: %v", r.Request.URL, err) // 记录请求错误信息
//	})
//
//	// 开始抓取目标网站
//	c.Visit("https://gorm.io/zh_CN/docs/")
//}
