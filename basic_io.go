package main // 如果是要可被執行的話，就要使用 main，有些是直接給其他程式或其他人使用時，就不需要 main
import "fmt" // 載入內建的 fmt 封包，用來做基本的輸入輸出

func Basic_io(){
	var x int 
	var y int
 fmt.Scanln(&x)
 fmt.Scanln(&y)
 fmt.Println(x * y)

}


// 數學運算符：+，-，*，/
// 賦值運算符：=，+=，-=，*=，/=
// 自增運算符：++，--
// 比較運算符：>，<，>=，<=，==
// 邏輯運算符：!，&&，||