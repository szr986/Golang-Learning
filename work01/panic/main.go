package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	collections = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, user := range users {
		each := strings.Split(user, "")
		for _, b := range each {
			if b == "e" || b == "E" {
				collections[user] += 1
				coins -= 1
			} else if b == "i" || b == "I" {
				collections[user] += 2
				coins -= 2
			} else if b == "o" || b == "O" {
				collections[user] += 3
				coins -= 3
			} else if b == "u" || b == "U" {
				collections[user] += 4
				coins -= 4
			}
		}
	}
	return coins
}

func main() {
	// left := dispatchCoin()
	// fmt.Println("剩下：", left)
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(collections)
}
