<!-- vscode-markdown-toc -->
* [New和Make的区别](#NewMake)
* 1. [list](#list)
	* 1.1. [list的底层结构](#list-1)
* 2. [slice](#slice)
	* 2.1. [slice的本质](#slice-1)
	* 2.2. [slice能否直接比较](#slice-1)
	* 2.3. [slice赋值拷贝](#slice-1)
	* 2.4. [slice的扩容策略](#slice-1)
* 3. [struct](#struct)
	* 3.1. [struct能否比较？](#struct-1)
* 4. [Defer](#Defer)
	* 4.1. [go defer的原理](#godefer)
	* 4.2. [defer执行时机](#defer)
* 5. [进程，线程，协程，锁与goroutine](#goroutine)
	* 5.1. [Goroutine](#Goroutine)
		* 5.1.1. [动态栈](#dongtaizhan)
		* 5.1.2. [Goroutine的调度](#Goroutine-1)
	* 5.2. [主协程如何等其余协程结束才操作？](#-1)
	* 5.3. [Golang的锁机制](#Golang)
* 6. [GO的垃圾回收(GC)](#GOGC)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc --># GolangLearning


记录golang学习过程中的重点和问题

##  1. <a name='NewMake'></a>New和Make的区别
1.make 只能初始化 slice、map和chan类型的对象，而new可以初始化任意类型  
2.make返回的是引用类型，而new返回的是指针类型  
3.make可以对三种类型（slice、map和chan）内部数据结构（长度和容量）进行初始化，而new不会


##  1. <a name='list'></a>list
###  1.1. <a name='list-1'></a>list的底层结构
在 Go 语言中，将列表使用 container/list 包来实现，内部的实现原理是双链表。列表能够高效地进行任意位置的元素插入和删除操作。

##  2. <a name='slice'></a>slice  
   

###  2.1. <a name='slice-1'></a>slice的本质
切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。   
![slice1](https://www.liwenzhou.com/images/Go/slice/slice_01.png)  
切片s2 := a[3:6]，相应示意图如下：  
![slice2](https://www.liwenzhou.com/images/Go/slice/slice_02.png)

###  2.2. <a name='slice-1'></a>slice能否直接比较  
切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一
合法的比较操作是和nil比较。 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都
是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：
```
var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
```
所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。  

###  2.3. <a name='slice-1'></a>slice赋值拷贝
下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。
```
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```

所以需要复制一个切片时，需要使用
`copy(s2,s1)`

###  2.4. <a name='slice-1'></a>slice的扩容策略



可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：

```
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

从上面的代码可以看出以下内容：  

1.首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。

2.否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）， 

3.否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for   {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）  

4.如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。 

##  3. <a name='struct'></a>struct
###  3.1. <a name='struct-1'></a>struct能否比较？
原文地址：[手撕 Go 面试官：Go 结构体是否可以比较，为什么？](https://blog.csdn.net/EDDYCJY/article/details/115327544?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1.pc_relevant_default&utm_relevant_index=1)


先给出结论：大部分不能，小部分能  
1. 当两个结构体的类型，字段完全相同时，可以比较  。
2. 其基本类型包含：slice、map、function 时，是不能比较的。

具体内容参见原文，写的很好

##  4. <a name='Defer'></a>Defer
###  4.1. <a name='godefer'></a>go defer的原理
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。  
例子：   
```
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
```
输出结果：  
```
start
end
3
2
1
```
###  4.2. <a name='defer'></a>defer执行时机
在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：  
![defer](https://www.liwenzhou.com/images/Go/func/defer.png)  

  
##  5. <a name='goroutine'></a>进程，线程，协程，锁与goroutine
回顾一下计算机操作系统的知识  

- 进程（process）：程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。

- 线程（thread）：操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位。

- 协程（coroutine）：非操作系统提供而是由用户自行创建和控制的用户态‘线程’，比线程更轻量级。  
  
Go语言中的并发程序主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现，当然也支持使用传统的多线程共享内存的并发方式。

###  5.1. <a name='Goroutine'></a>Goroutine
Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。例如Go运行时会智能地将 m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。

Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine——main goroutine，当 Go 程序启动时它会自动创建。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。
####  5.1.1. <a name='dongtaizhan'></a>动态栈
操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。
####  5.1.2. <a name='Goroutine-1'></a>Goroutine的调度
操作系统的线程会被操作系统内核调度时会挂起当前执行的线程并将它的寄存器内容保存到内存中，选出下一次要执行的线程并从内存中恢复该线程的寄存器信息，然后恢复执行该线程的现场并开始执行线程。从一个线程切换到另一个线程需要完整的上下文切换。因为可能需要多次内存访问，索引这个切换上下文的操作开销较大，会增加运行的cpu周期。

区别于操作系统内核调度操作系统线程，goroutine 的调度是Go语言运行时（runtime）层面的实现，是完全由 Go 语言本身实现的一套调度系统——go scheduler。它的作用是按照一定的规则将所有的 goroutine 调度到操作系统线程上执行。

在经历数个版本的迭代之后，目前 Go 语言的调度器采用的是 GPM 调度模型。  
![gpm]https://www.liwenzhou.com/images/Go/concurrence/gpm.png

其中：

- G：表示 goroutine，每执行一次go f()就创建一个 G，包含要执行的函数和上下文信息。

全局队列（Global Queue）：存放等待运行的 G。

- P：表示 goroutine 执行所需的资源，最多有 GOMAXPROCS 个。

- P 的本地队列：同全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建 G 时，G 优先加入到 P 的本地队列，如果本地队列满了会批量移动部分 G 到全局队列。

- M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，当 P 的本地队列为空时，M 也会尝试从全局队列或其他 P 的本地队列获取 G。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。

Goroutine 调度器和操作系统调度器是通过 M 结合起来的，每个 M 都代表了1个内核线程，操作系统调度器负责把内核线程分配到 CPU 的核上执行。


###  5.2. <a name='-1'></a>主协程如何等其余协程结束才操作？
1.共享内存  
设置一个全局变量，子协程对全局变量进行修改，当该变量等于某一值时主协程才结束。  

这样我们是可以做到等待其他协程执行结束的。但是并不是 Go 所提倡的：
```
Do not communicate by sharing memory; instead, share memory by communicating.

不要以共享内存的方式来通信，相反，要通过通信来共享内存。
```


2.time.sleep(傻逼方法···)

3.WaitGroup
```
var wg sync.WaitGroup

func f() {
	rand.Seed(time.Now().UnixNano()) //生成随机数
	defer wg.Done()
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func main() {
	wg.Add(1)
	go f()
	wg.Wait()
}
```

4.channel  
使用一个无缓冲区的channel，主协程从channel中获取值，无法从中取得值时主协程会阻塞
```
var b chan int

func f() {
	rand.Seed(time.Now().UnixNano()) //生成随机数

	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
	b <- 1
}

func main() {
	b = make(chan int)
	go f()
	<-b
	return
}
```

###  5.3. <a name='Golang'></a>Golang的锁机制
Golang 中的有两种锁，为 sync.Mutex 和 sync.RWMutex。  

- sync.Mutex 互斥锁，只有一种锁：Lock()，它是绝对锁，同一时间只能有一个锁。
- sync.RWMutex 读写锁，它有两种锁： RLock() 和 Lock()：
RLock() 叫读锁。它不是绝对锁，比起互斥锁有着更高的并行性，它允许多个读者同时读取。
Lock() 叫写锁，它是个绝对锁，就是说，如果一旦某人拿到了这个锁，别人就不能再获取此锁了。

总结  
- 正常情况下，在请求 Lock() 锁时发现资源被锁住了，无论是 RLock() 锁还是 Lock() 锁，它都会等待。
- 正常情况下，在请求 RLock() 锁时发现资源被 Lock() 锁住了，它会等待。发现是被 RLock()
锁住，自己也可以读取。（这个是用数字的原子操作来控制的，原理见附的文章的源码解释）
- 不要嵌套地去用 锁，这样则有可能发生死锁，即大家（所有 goroutine）都在等待锁的释放，此时发生死锁。

##  6. <a name='GOGC'></a>GO的垃圾回收(GC)
内容过多，详见文章：[Go面试题(六)：一文弄懂 Golang GC、三色标记、混合写屏障机制【图文解析GC】](https://blog.csdn.net/xiaodaoge_it/article/details/121890145)
