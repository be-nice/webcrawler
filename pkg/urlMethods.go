package pkg

import (
	"errors"
	"net/url"
	"path"
	"regexp"
)

func NormalizeURL(s string) (string, error) {
	re := regexp.MustCompile(`^https?://`)
	s = re.ReplaceAllString(s, "")
	data, err := url.Parse(s)
	if err != nil {
		return "", errors.New("Weird error")
	}

	data.Path = path.Clean(data.Path)

	return data.String(), nil
}
