package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	value := `[{"path":"D:/xxx/mysql.log","topic":"web_log"},{"path":"D:/xxx/redis.log","topic":"web_log"}]
	`
	// value := `[{"path":"c:/tmp/nginx.log","topic":"nginx_log"},{"path":"D:/xxx/redis.log","topic":"redis_log"},{"path":"D:/xxx/mysql.log","topic":"mysql_log"}]
	// `
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/192.168.0.100/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

}
