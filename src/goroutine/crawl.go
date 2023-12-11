package main

import (
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func StartCrawl(maxValGoroutine int) {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < maxValGoroutine; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
