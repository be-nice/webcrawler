package pkg

import "sync"

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
