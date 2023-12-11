/*
缓存（memoizing）函数，需要将缓存函数的返回结果，这样在对函数进行调用的时候只要返回计算的结果就可以了，不需要再次进行计算。

解决方案：并发安全且会避免对整个缓存加锁而导致所有操作都去竞争一个锁的设计
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/lutianen/gopl/src/memoizing/memo"
)

func main() {
	m := memo.New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			strat := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(strat), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func httpGetBody(url string) (any, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",

		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}
}
