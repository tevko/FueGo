package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/html"
)

// Worker function to fetch and scrape a single page
func worker(id int, urlChan <-chan string, results chan<- map[string]string, wg *sync.WaitGroup) error {
	defer wg.Done()
	for u := range urlChan {
		fmt.Printf("Worker %d fetching: %s\n", id, u)
		resp, err := http.Get(u)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			return err
		}
		l, err := url.Parse(u)
		if err != nil {
			return err
		}
		data := extractLinks(doc, l)
		results <- data
	}
	return nil
}

func extractLinks(doc *html.Node, baseURL *url.URL) map[string]string { 
	data := make(map[string]string)
	var extractLinksRecursive func(*html.Node)
	extractLinksRecursive = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
					for _, a := range n.Attr {
							if a.Key == "href" {
									link, err := baseURL.Parse(a.Val)
									if err != nil {
											fmt.Print(err)
											continue 
									}
									data[link.String()] = link.String()
									break
							}
					}
			} else {
					for c := n.FirstChild; c != nil; c = c.NextSibling {
							extractLinksRecursive(c)
					}
			}
	}
	extractLinksRecursive(doc)
	return data
}

func main() {
	var wg sync.WaitGroup
	// URLs to scrape
	urls := []string{
		"https://google.com",
		"https://old.reddit.com/",
		"https://timevko.website",
	}
	urlChan := make(chan string)
	results := make(chan map[string]string)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, urlChan, results, &wg)
	}
	for _, url := range urls {
		urlChan <- url
	}
		close(urlChan)
    go func() {
			wg.Wait()
			close(results)
		}()

	// Collect results (receive until the channel is closed)
	for result := range results {
		fmt.Println("Extracted links:")
		for link := range result {
			fmt.Println(link)
		}
	}
}
