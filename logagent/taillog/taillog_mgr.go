package taillog

import (
	"fmt"
	"time"
	"work/logagent/etcd"
)

var tskMgr *tailLogMgr

// tailtask管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲通道
	}
	for _, logEntry := range logEntryConf {
		// 启动tailtask
		tailtask := NewTailTask(logEntry.Path, logEntry.Topic)
		// 记录启动了哪些tailtask
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailtask

	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有新的配置过来之后做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				// 将path和topic拼接，形成唯一的识别
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]

				fmt.Printf("func run newconf比较，mk:%v,是否存在:%v\n", mk, ok)
				// 判断是否有新增
				if ok {
					// 原来有就不需要操作
					continue
				} else {
					// 如果没有 新增一个对象
					tailObj := NewTailTask(conf.Path, conf.Topic)
					fmt.Println("新tailObj，Path：", tailObj.path)
					t.tskMap[mk] = tailObj
				}
			}

			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { //从原来的配置中依次拿出配置项，去新的配置中逐一进行比较
				fmt.Println("比较新老配置：", c1.Path)
				isDelete := true
				for _, c2 := range newConf {
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						fmt.Println("Path:%v,Topic:%v", c2.Path == c1.Path, c2.Topic == c1.Topic)
						isDelete = false
						continue
					}
				}
				if isDelete {
					fmt.Println("删除:", c1.Path)
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
					// delete(t.tskMap, mk)
					// t.logEntry = newConf
					// fmt.Println(t.tskMap)
				}
			}
			fmt.Println("新配置:", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外暴露，tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
