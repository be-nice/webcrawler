package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func GetHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error: status code %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	normCurr, err := NormalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawBaseURL, rawCurrentURL)

	if _, ok := pages[normCurr]; ok {
		pages[normCurr]++
		return
	}

	pages[normCurr] = 1
	html, _ := GetHTML(rawCurrentURL)
	data, _ := ScanPageForURL(html, rawBaseURL)

	u1, _ := url.Parse(rawBaseURL)
	for _, val := range data {
		u2, _ := url.Parse(val)

		if u1.Hostname() != u2.Hostname() {
			continue
		}
		CrawlPage(rawBaseURL, val, pages)
	}
}

func (c *Crawler) CrawlDomain() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	workQueue := make(chan string, 1000)

	baseParsed, err := url.Parse(c.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	crawl := func() {
		for rawURL := range workQueue {
			normURL, err := NormalizeURL(rawURL)
			if err != nil {
				log.Printf("normalize error: %v\n", err)
				wg.Done()
				continue
			}

			mu.Lock()
			if _, ok := c.Pages[normURL]; ok {
				c.Pages[normURL]++
				mu.Unlock()
				wg.Done()
				continue
			}

			if len(c.Pages) >= c.MaxPages {
				mu.Unlock()
				wg.Done()
				continue
			}

			c.Pages[normURL] = 1
			mu.Unlock()

			html, err := GetHTML(rawURL)
			if err != nil {
				log.Printf("fetch error: %v\n", err)
				wg.Done()
				continue
			}

			links, err := ScanPageForURL(html, c.BaseURL)
			if err != nil {
				log.Printf("scan error: %v\n", err)
				wg.Done()
				continue
			}

			for _, link := range links {
				linkParsed, err := url.Parse(link)
				if err != nil {
					continue
				}
				if linkParsed.Hostname() == baseParsed.Hostname() {
					mu.Lock()
					if len(c.Pages) >= c.MaxPages {
						mu.Unlock()
						break
					}
					mu.Unlock()
					wg.Add(1)
					workQueue <- link
				}
			}

			wg.Done()
		}
	}

	for range c.MaxWorkers {
		go crawl()
	}

	wg.Add(1)
	workQueue <- c.BaseURL

	wg.Wait()
	close(workQueue)
	for range workQueue {
		// Drainage
	}
}
