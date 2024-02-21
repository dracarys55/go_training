package main // 如果是要可被執行的話，就要使用 main，有些是直接給其他程式或其他人使用時，就不需要 main
import "fmt" // 載入內建的 fmt 封包，用來做基本的輸入輸出

func Data_var(){
 fmt.Println(3) //整數 int
 fmt.Println(3.1415) // 浮點數 float 64
 fmt.Println("測試中文") //字串 string
 fmt.Println(true) //布林值 bool 
 fmt.Println('a') //字符 rune 

 // 也可以這樣操作
 fmt.Println(3,3,3,3,3,3,3)

}