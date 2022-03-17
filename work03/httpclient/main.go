package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Println("get url failed :", err)
	// 	return
	// }
	data := url.Values{} //url encode
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周丽")
	data.Set("age", "9000")
	queryStr := data.Encode() //URL encode之后的URL
	fmt.Println(queryStr)
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	// 发请求
	resq, err := http.DefaultClient.Do(req)
	// 从resp中把服务端返回的数据读出来
	b, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		fmt.Println("read resp.Body failed:", err)
		return
	}
	fmt.Println(string(b))

}
