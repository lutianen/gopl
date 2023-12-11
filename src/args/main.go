package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 每次循环迭代字符串 s 的内容都会更新，s 原来的内容已经不再使用，将在适当时机对它进行垃圾回收
	// 如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用 strings 包的 Join 函数
	/*
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	*/

	fmt.Println(strings.Join(os.Args[1:], " "))
}
