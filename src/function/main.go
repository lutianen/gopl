package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "herf" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// findLinks performs an HTTP GET request for url, parses the response as HTML, and extracts and returns the links
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func sum(vs ...int) int {
	total := 0
	for _, v := range vs {
		total += v
	}
	return total
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findLinks: %c\n", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}

	fmt.Println("-----------")
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}
