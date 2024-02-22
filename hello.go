package main // 如果是要可被執行的話，就要使用 main，有些是直接給其他程式或其他人使用時，就不需要 main
// 載入內建的 fmt 封包，用來做基本的輸入輸出 fmt 是格式化（formatting）的縮寫。

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
		log.Fatal("404")
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello!")
}


func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

    /* 
    if err := http.ListenAndServe(":8080", nil); err != nil {

    http.ListenAndServe 函數會嘗試在指定的地址和端口上啟動一個 HTTP 服務器。在這個例子中，地址是空字符串（代表所有網絡接口），端口是 8080。

    第二個參數是 http.Handler，這裡傳入 nil 表示使用預設的處理器，也就是 http.DefaultServeMux。

    := 是短變數聲明，用於在 if 語句的範圍內創建並初始化 err 變數。這個變數會接收 http.ListenAndServe 返回的錯誤值。

    err != nil {

    這是 if 條件語句，用於檢查 err 變數是否不為 nil。在 Go 中，如果一個函數返回的錯誤值（error）不為 nil，則表示有錯誤發生。
    log.Fatal(err)

    如果錯誤值不為 nil，則調用 log.Fatal 函數。log.Fatal 會輸出錯誤訊息到標準錯誤輸出，然後調用 os.Exit(1) 退出程序，退出碼 1 通常表示程序因為錯誤而終止。
    }

    這是 if 語句的結束大括號
        

    在 Go 語言中，nil 是一個預定義的識別符號，用於表示某些類型的零值（zero value）。具體來說，它用於表示指針、通道（channel）、函數、接口、映射（map）或切片（slice）的零值。如果這些類型的變數沒有初始化為具體的值，它們的預設值就是 nil。這表示變數沒有指向任何有效的對象、數據結構、函數等。
在 JavaScript 中，null 用來表示一個故意設置的空或不存在的值。它是 JavaScript 原始值之一，用於表示空值或不存在的對象引用。

雖然 Go 的 nil 和 JavaScript 的 null 在概念上有些相似，都表示“無”或“空”的狀態，但是它們在兩種語言中的具體用法和含義是不同的。
在 Go 中，nil 有更多的使用場景（如上所述），而且 Go 是靜態類型的，這意味著只有特定的類型才能賦值為 nil。而在 JavaScript 中，null 是一個可以賦值給任何變數的原始值
它和 undefined 一起，用於表示缺失的值。
    */
}


/* func main(){
	fmt.Println("Hello world")
	fmt.Println('a')
	 // Data_var() 公開函數以大寫命名
	 // Basic_io()
}
*/

// 了解 go mod （類似 node moudles


// Print

// Print 函數將它的參數連接成一個字符串，然後輸出到標準輸出（通常是終端或命令行界面）。
// 參數之間不會自動添加空格。
// 輸出後不會添加換行符，所以如果你連續調用 Print，輸出會在同一行上連接起來。
// Printf

// Printf 函數提供了格式化輸出。你可以使用格式化的動詞（如 %s、%d、%v 等）來指定每個參數應該如何顯示。
// 這非常類似於 C 語言中的 printf 函數。
// 使用格式化字符串作為第一個參數，後面跟著一個或多個參數，這些參數將替換格式化字符串中對應的格式化動詞。
// Println
// ln = line 的意思

// Println 函數也將參數連接成一個字符串，但它會在參數之間自動插入空格，並在輸出結束後添加一個換行符。
// 這使得 Println 非常適合快速列印值，同時保持它們之間的分隔，並在每次調用後開始新的一行。

// 總之，你會選擇 Print 當你想有更精細的控制輸出但不需要格式化的時候；
// 選擇 Printf 當你需要根據特定的格式模板輸出文字的時候；而 Println 適用於快速、簡單地列印信息，且每次調用後都會開始新行。