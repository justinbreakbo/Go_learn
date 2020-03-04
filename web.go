package main

import (
	"fmt"
	"net/http"
)
// 处理器函数
func Read(w http.ResponseWriter, r *http.Request) {
	// w 当请求传送过来的时候,这个w 能够返回数据给请求
	_, _ = w.Write([]byte("Hello Web"))
}

func Read1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	_, _ = w.Write([]byte("Hello Web1"))
}


func main() {
	// 处理器
	mux := http.NewServeMux()

	mux.HandleFunc("/",Read)
	mux.HandleFunc("/test",Read1)

	err := http.ListenAndServe(":9800",mux)
	if err != nil {
		fmt.Println(err)
	}
}