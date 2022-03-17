package taillog

import (
	"context"
	"fmt"
	"work/logagent/kafka"

	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

// 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了能够实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.Init()
	return
}

func (t *TailTask) Init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,                                 //
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed:", err)
	}

	// 当goroutine执行的函数退出时，goroutine结束
	go t.run()
}

// // 专门读取日志
// func Init(filename string) (err error) {
// 	fileName := "./my.log"
// 	config := tail.Config{
// 		ReOpen:    true,                                 //重新打开
// 		Follow:    true,                                 //是否跟随
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪个地方开始读
// 		MustExist: false,                                //文件不存在不报错
// 		Poll:      true,                                 //
// 	}
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("tail file failed:", err)
// 		return
// 	}
// 	return
// }

// func (t *TailTask) ReadChan() <-chan *tail.Line {
// 	return t.instance.Lines
// }

// func ReadLog() <-chan *tail.Line {
// 	// var (
// 	// 	line *tail.Line
// 	// 	ok   bool
// 	// )
// 	// for {
// 	// 	line, ok = <-tailObj.Lines
// 	// 	if !ok {
// 	// 		fmt.Printf("tail file close reopen, filename:%s\n", tailObj.Filename)
// 	// 		time.Sleep(time.Second)
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println("line:", line.Text)
// 	return tailObj.Lines
// }

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("Tail task:%v exit\n", t.path+t.topic)
			return
		case line := <-t.instance.Lines:
			// 先把日志数据发送到通道中
			// kafka中有专门的函数去取通道再发送
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
