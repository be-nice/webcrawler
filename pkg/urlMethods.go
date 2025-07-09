package pkg

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func NormalizeURL(s string) (string, error) {
	// re := regexp.MustCompile(`^https?://`)
	// s = re.ReplaceAllString(s, "")
	data, err := url.Parse(s)
	if err != nil {
		return "", errors.New("Weird error")
	}

	path := strings.TrimRight(data.EscapedPath(), "/")

	if path == "" || path == "/" {
		return data.Hostname(), nil
	}

	return data.Hostname() + path, nil
}

func ScanPageForURL(rawBody, rootURL string) ([]string, error) {
	links := []string{}
	base, err := url.Parse(rootURL)
	if err != nil {
		return links, err
	}

	doc, err := html.Parse(strings.NewReader(rawBody))
	if err != nil {
		return links, err
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n == nil {
			return
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val
					parsedHref, err := url.Parse(href)
					if err == nil {
						abs := base.ResolveReference(parsedHref)
						links = append(links, abs.String())
					}
				}
			}
		}

		traverse(n.FirstChild)
		traverse(n.NextSibling)
	}

	traverse(doc)

	return links, nil
}
