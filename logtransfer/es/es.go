package es

import (
	"context"
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

// 初始化es，准备接收kafka发来的数据

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connect to es success")
	return
}

// 发送数据到es
func SendToES(indexStr string, data interface{}) error {
	fmt.Println(indexStr, data)
	put1, err := client.Index().
		Index(indexStr).
		BodyJson(data).
		// Type("xxx").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return err
}
