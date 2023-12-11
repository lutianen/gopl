package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		// 以 流 的模式读取输入
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			counts[input.Text()]++
		}

		// NOTE: ignore potential errors from input.Err()
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	// 一口气把全部全部输入数据读入到内存中，一次分割为多行
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// data, err := ioutil.ReadFile(filename)
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			continue
		}

		// strings.Split 与 strings.Join 功能相反
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
		}
	}
}
