package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

// 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化etcd
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	return
}

// 从etcd中根据key获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed:", err)
			return
		}
	}
	// fmt.Println(logEntryConf)
	return
}

//watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// 通知taillog_mgr
			// 先判断操作类型
			var newConf []*LogEntry
			if ev.Type != clientv3.EventTypeDelete {
				// 如果是删除操作，手动传递一个空的配置项，否则unmarshal会报错
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("unmarshal error:", err)
					continue
				}
			}
			fmt.Println("get new conf :", newConf)
			newConfCh <- newConf
		}
	}
}
