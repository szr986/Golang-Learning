package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server
func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于Get请求，参数都在URL上，请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL的query param
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
