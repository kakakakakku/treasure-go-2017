package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type Page struct {
	Title       string
	Description string
}

func Get(url string) (*Page, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// https://godoc.org/golang.org/x/net/html
	docs, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := &Page{}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			p.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			if isDescription(n.Attr) {
				for _, attr := range n.Attr {
					p.Description = attr.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(docs)

	return p, err
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description" {
			return true
		}
	}
	return false
}

func main() {
	p, err := Get("https://voyagegroup.com/")
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("%#v", p)
}
