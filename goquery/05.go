package main

// goquery选择器
// 类似jQuery 或者css选择器，常用的有元素名称选择器、ID选择器、class选择器

//// 元素名称选择器
//func main() {
//	html := `<body>
//		<div>DIV1</div>
//		<div>DIV2</div>
//		<span>SPAN</span>
//	</body>`
//
//	// 从HTML字符串创建一个新的Document对象
//	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// 使用元素名称选择器选择所有的<div>元素
//	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(i, "select text", selection.Text()) // 打印每个<div>元素的文本内容
//	})
//}

//// ID选择器
//func main() {
//	html := `<body>
//		<div id="div1">DIV1</div>
//		<div>DIV2</div>
//		<span>SPAN</span>
//	</body>`
//
//	// 从HTML字符串创建一个新的Document对象
//	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// 使用ID选择器选择ID为"div1"的元素
//	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(selection.Text()) // 打印ID为"div1"的元素的文本内容
//	})
//}

//// class类选择器
//func main() {
//	html := `<body>
//		<div id="div1">DIV1</div>
//		<div class="name">DIV2</div>
//		<span>SPAN</span>
//	</body>`
//
//	// 从HTML字符串创建一个新的Document对象
//	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	// 使用Class选择器选择class为"name"的元素
//	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(selection.Text()) // 打印class为"name"的元素的文本内容
//	})
//}

//// goquery常用方法
//// Selection 可以提取，删除，添加元素，属性内容
//// 1)位置操作函数
//// Eq(index int) *Selection //根据索引获取某个节点集
//selection.Eq(2) // 获取索引位置为2的节点
//// First() *Selection //获取第一个子节点集
//selection.First() // 获取第一个子节点
//// Last() *Selection //获取最后一个子节点集
//selection.Last() // 获取最后一个子节点
//// Next() *Selection //获取下一个兄弟节点集
//selection.Next() // 获取下一个兄弟节点
//// NextAll() *Selection //获取后面所有兄弟节点集
//selection.NextAll() // 获取后面所有兄弟节点
//// Prev() *Selection //前一个兄弟节点集
//selection.Prev() // 获取前一个兄弟节点
//// Get(index int) *html.Node //根据索引获取一个节点
//selection.Get(2) // 根据索引获取节点
//// Index() int //返回选择对象中第一个元素的位置
//selection.Index() // 返回选择对象中第一个元素的位置
//// Slice(start, end int) *Selection //根据起始位置获取子节点集
//selection.Slice(0, 2) // 根据起始位置获取子节点集
//// 2)循环遍历选择的节点
//// Each(f func(int,*Selection)) *Selection //遍历
//// EachWithBreak(f func(int, *Selection)bool) *Selection //可中断遍历
//// Map(f func(int,*Selection)string)(result[]string) //返回字符串数组
//// 3)检测或获取节点属性值
//// Attr()，RemoveAttr()，SetAttr()
//value, exists := selection.Attr("href") // 获取属性值
//selection.RemoveAttr("href") // 移除属性
//selection.SetAttr("href", "https://example.com") // 设置属性值
//// Addclass(),Hasclass(),Removeclass(),Toggleclass()
//selection.AddClass("new-class") // 添加class
//hasClass := selection.HasClass("existing-class") // 检查是否有class
//selection.RemoveClass("old-class") // 移除class
//selection.ToggleClass("toggle-class") // 切换class
//// Html() //获取该节点的html
//html, err := selection.Html()
//// Length() //返回该Selection的元素个数
//length := selection.Length()
//// Text() //获取该节点的文本值
//text := selection.Text()
//// 4)在文档树之间来回跳转(常用的查找节点方法)
//// Children() //返回selection中各个节点下的孩子节点
//children := selection.Children()
//// Contents() //获取当前节点下的所有节点
//contents := selection.Contents()
//// Find() //查找获取当前匹配的元素
//found := selection.Find(".some-class")
//// Next() //下一个元素
//next := selection.Next()
//// Prev() //上一个元素
//prev := selection.Prev()
