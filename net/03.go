package main

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
//	req.Header.Add("Cookie", "_ga=GA1.1.1901661351.1706858808; _ga_YXBYDX14GJ=GS1.1.1721139829.15.1.1721139958.58.0.0")
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
//
//	save(title, content)
//}
//
////// 保存在本地文件
////func save(title string, content string) {
////	err := os.WriteFile("./article/"+title+".html", []byte(content), 0644)
////	if err != nil {
////		panic(err)
////	}
////}
//
//var DB *gorm.DB
//
//func MysqlInit() { // 初始化数据库
//
//	user := ""
//	pass := ""
//	port := ""
//	host := ""
//	name := "golang_reptile"
//
//	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		DisableForeignKeyConstraintWhenMigrating: true,
//	})
//
//	if err != nil {
//		log.Fatal("DatabaseConnectFailed", err)
//	}
//
//	err = autoMigrate(db)
//	if err != nil {
//		log.Fatal("DatabaseMigrateFailed", err)
//	}
//	DB = db
//}
//
//type reptile struct {
//	ID          int
//	Title       string
//	Content     string `gorm:"text"`
//	CreatedTime time.Time
//}
//
//func autoMigrate(db *gorm.DB) error {
//	return db.AutoMigrate(
//		&reptile{},
//	)
//}
//
//func save(title string, content string) {
//	if DB == nil {
//		MysqlInit()
//	}
//
//	rept := reptile{
//		Title:       title,
//		Content:     content,
//		CreatedTime: time.Now(),
//	}
//	result := DB.Create(&rept)
//	if result.Error != nil {
//		log.Printf("Save to database failed: %v", result.Error)
//	} else {
//		log.Println("Saved to database successfully")
//	}
//}
//
//func main() {
//	url := "https://gorm.io/zh_CN/docs/"
//	s, err := request(url)
//	if err != nil {
//		log.Fatalf("Failed to fetch URL: %v", err)
//	}
//	parse(s)
//}
