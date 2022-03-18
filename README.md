# Golang-Learning
记录下golang学习过程中的难点和关键点
### New和Make的区别
1.make 只能初始化 slice、map和chan类型的对象，而new可以初始化任意类型  
2.make返回的是引用类型，而new返回的是指针类型
3.make可以对三种类型（slice、map和chan）内部数据结构（长度和容量）进行初始化，而new不会
