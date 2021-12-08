package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func m1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
	    fmt.Fprintf(os.Stderr, "err:%v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string{
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling{
		links = visit(links, c)  // 递归
	}
	return links
}