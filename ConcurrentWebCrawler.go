package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

type Link struct {
	URL string
}

func main() {
	links := []string{
		"https://github.com/",
		"https://chatgpt.com/",
		"https://sabzlearn.ir/",
		"https://fastdic.com/",
		"https://www.youtube.com/",
		"http://google.com/",
		"https://www.arefonline.ir/",
	}
	linksCount := len(links)
	workersCount := 3
	done := make(chan bool, linksCount)
	linkChan := make(chan Link, linksCount)
	for i := 0; i < workersCount; i++ {
		go worker(i, linkChan, done)
	}

	for _, value := range links {
		linkChan <- Link{URL: value}
	}
	close(linkChan)

	for i := 0; i < workersCount; i++ {
		<-done
	}
}

func worker(workerId int, linkChan <-chan Link, done chan<- bool) {
	fmt.Printf("worker %d starts \n", workerId)
	for link := range linkChan {
		fmt.Printf("Worker %d starts crawling %s \n", workerId, link.URL)
		res, err := http.Get(link.URL)
		if err != nil {
			fmt.Println("error in get", err.Error())
			continue
		}

		doc, err := html.Parse(res.Body)
		if err != nil {
			fmt.Println("error in parse", err.Error())
			continue
		}

		title := getTitle(doc)
		
		fmt.Printf("Worker %d ends crawling %s and result is %s \n", workerId, link.URL, title)
	}
	done <- true

}
func getTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		result := getTitle(child)
		if result != "" {
			return result
		}
	}
	return ""
}
