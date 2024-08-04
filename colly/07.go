package main

//func main() {
//	// 创建一个新的 Colly 收集器
//	c := colly.NewCollector()
//
//	// 在每个匹配到的 .sidebar-link 元素上执行的回调函数
//	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
//		// 访问每个链接的 href 属性所指向的 URL
//		e.Request.Visit(e.Attr("href"))
//	})
//
//	// 在每个请求发出之前执行的回调函数
//	c.OnRequest(func(r *colly.Request) {
//		// 打印即将访问的 URL
//		fmt.Println("Visiting URL:", r.URL)
//	})
//
//	// 开始访问目标网站
//	c.Visit("https://gorm.io/zh_CN/docs/")
//}

// 回调方法
// OnRequest 在请求之前调用
// OnError  在请求期间发生错误时调用
// OnResponseHeaders 在收到响应标头后调用
// OnResponse 收到响应后调用
// OnHTML 如果接收到的内容是 HTML，则立即调用
// OnXML
// OnHTML 如果接收到的内容是 HTML或 XML，则立即调用
// OnScraped 回调后 OnXML 调用

//func main() {
//	c := colly.NewCollector()
//	c.OnRequest(func(r *colly.Request) {
//		fmt.Println("请求前调用:OnRequest")
//	})
//
//	c.OnError(func(_ *colly.Response, err error) {
//		fmt.Println("发生错误调用:OnError")
//	})
//
//	c.OnResponse(func(r *colly.Response) {
//		fmt.Println("获得响应后调用:OnResponse")
//	})
//
//	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//		fmt.Println("OnResponse收到html内容后调用:OnHTML")
//	})
//
//	c.OnXML("//h1", func(e *colly.XMLElement) {
//		fmt.Println("OnResponse收到xmL内容后调用:OnXML")
//	})
//
//	c.OnScraped(func(r *colly.Response) {
//		fmt.Println("结束", r.Request.URL)
//	})
//
//	c.Visit("https://gorm.io/zh_CN/docs/")
//}
