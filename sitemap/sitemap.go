package sitemap

import (
	"fmt"
	"net/http"

	"github.com/Shamanskiy/gophercises/linkparser"
)

func ParseSite(url string) ([]string, error) {
	url += "/x"
	fmt.Printf("Client: GET %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	links, err := linkparser.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return getHRefs(links), nil
}

func getHRefs(links []linkparser.Link) []string {
	hrefs := make([]string, len(links))
	for i, link := range links {
		hrefs[i] = link.Href
	}
	return hrefs
}
