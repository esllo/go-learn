package article

var articles = []func(){Article1_Print, Article2_FormattedPrint, Article3_Variables, Article4_DataType, Article5_Conditional, Article6_Loop, Article_Function, Article_AnonymousFunction, Article_Closure, Article_Array, Article_Map, Article_Struct, Article_Method, Article_Interface, Article_Error, Article_Panic, Article_Goroutine, Article_Channel}

func PrintArticle(i int) {
	if i < 0 || i > len(articles)-1 {
		return
	}
	println("================")
	println("article", i)
	println("----------------")
	articles[i]()
	println("\n================\n\n")
}

func GetArticleLength() int {
	return len(articles)
}

func PrintAllArticles() {
	for index := range articles {
		PrintArticle(index)
	}
}
