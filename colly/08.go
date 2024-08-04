package main

//// colly的配置
//// 设置UserAgent
//c := colly.NewCollector()
//c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.0.0"
//// 设置cookie
//siteCookie := c.Cookies("url")
//c.SetCookies("", siteCookie)

//// HTTP配置
//c:= colly.NewCollector()
//c.WithTransport(&http.Transport{
//	Proxy: http.ProxyFromEnvironment,
//	DialContext: (&net.Dialer{
//		Timeout:   30 * time.Second,
//		KeepAlive: 30 * time.Second,
//		DualStack: true,
//	}).DialContext,
//	MaxIdleConns:          100,
//	IdleConnTimeout:       90 * time.Second,
//	TLSHandshakeTimeout:   10 * time.Second,
//	ExpectContinueTimeout: 1 * time.Second,
//}
//}

//// colly页面爬取和解析
//func main() {
//	c := colly.NewCollector()
//	// selector goquery name id class
//	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
//		//e.Request.Visit(e.Attr("href"))
//		//// "#sidebar"
//		//fmt.Println("e.Text: %v\n", e.Text)
//		//fmt.Println("e.Name: %v\n", e.Name)
//		//r, _ := e.DOM.Html()
//		//fmt.Println("ret: %v\n", r)
//		// 绝对地址
//		s := e.Attr("href")
//		//fmt.Println("s: %v\n", s)
//		s2 := e.Request.AbsoluteURL(s)
//		fmt.Println("s2: %v\n", s2)
//		c.Visit(s2)
//	})
//	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
//		s := h.Text
//		fmt.Println("s: %v\n", s)
//	})
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("url", r.URL)
//	})
//	c.Visit("https://gorm.io/zh_CN/docs/")
//}

//func main() {
//	// 创建一个新的 Colly 收集器，并禁止重复访问相同的 URL
//	c := colly.NewCollector(
//		colly.MaxDepth(2), // 限制抓取深度为2
//	)
//
//	// 设置最大并发请求数量为2，每个请求之间的延迟为2秒
//	c.Limit(&colly.LimitRule{
//		DomainGlob:  "*",
//		Parallelism: 2,
//		Delay:       2 * time.Second,
//	})
//
//	// 定义在每个 .sidebar-link 元素上执行的回调函数
//	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
//		// 获取链接的 href 属性并转换为绝对 URL
//		link := e.Request.AbsoluteURL(e.Attr("href"))
//		fmt.Println("Found link:", link)
//		// 访问该链接
//		c.Visit(link)
//	})
//
//	// 定义在每个 .article-title 元素上执行的回调函数
//	c.OnHTML(".article-title", func(h *colly.HTMLElement) {
//		title := h.Text
//		fmt.Println("Article title:", title)
//	})
//
//	// 定义在每个请求发出之前执行的回调函数
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("Visiting URL:", r.URL)
//	})
//
//	// 定义错误处理回调函数
//	c.OnError(func(r *colly.Response, err error) {
//		log.Printf("Request URL: %v failed with response: %v\nError: %v", r.Request.URL, r, err)
//	})
//
//	// 开始抓取目标网站
//	c.Visit("https://gorm.io/zh_CN/docs/")
//}
