package pkg

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

func (c *Crawler) handlePageVisit(mu *sync.Mutex, link string) bool {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := c.Pages[link]; ok {
		c.Pages[link]++
		return false
	}

	if len(c.Pages) >= c.MaxPages {
		return false
	}

	c.Pages[link] = 1

	return true
}

func (c *Crawler) handleWorkQueue(mu *sync.Mutex, link string, base *url.URL) bool {
	linkParsed, err := url.Parse(link)
	if err != nil {
		return false
	}

	if linkParsed.Hostname() != base.Hostname() {
		return false
	}

	mu.Lock()
	defer mu.Unlock()
	if len(c.Pages) >= c.MaxPages {
		return false
	}

	return true
}

func ValidateScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Sprintf("http://%s", url)
	}

	return url
}
