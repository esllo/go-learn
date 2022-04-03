package main

import (
	"go-learn/article"
)

func printArticle(i int, f func()) {
	println("================")
	println("article", i)
	println("----------------")
	f()
	println("\n================\n\n")
}

func main() {
	printArticle(1, article.Article1_Print)
	printArticle(2, article.Article2_FormattedPrint)
	printArticle(3, article.Article3_Variables)
	printArticle(4, article.Article4_DataType)
	printArticle(5, article.Article5_Conditional)
	printArticle(6, article.Article6_Loop)
	printArticle(7, article.Article_Function)
	printArticle(8, article.Article_AnonymousFunction)
	printArticle(9, article.Article_Closure)
	printArticle(10, article.Article_Array)
	printArticle(11, article.Article_Map)
	printArticle(12, article.Article_Struct)
	printArticle(13, article.Article_Method)
	printArticle(14, article.Article_Interface)
	printArticle(15, article.Article_Error)
}
