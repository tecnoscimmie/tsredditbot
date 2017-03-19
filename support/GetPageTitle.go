package support

import (
	"errors"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

// GetPageTitle returns the <title> of a given page
func GetPageTitle(pageURL string) (string, error) {
	r, err := http.Get(pageURL)
	if err != nil {
		log.Fatal(err)
	}

	title := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	p, _ := html.Parse(r.Body)
	f(p)

	if title == "" {
		return "", errors.New("cannot find page title")
	}

	return title, nil
}
