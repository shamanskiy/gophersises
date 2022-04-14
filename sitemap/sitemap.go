package sitemap

import (
	"log"
	"net/http"
	"strings"

	"github.com/Shamanskiy/gophercises/linkparser"
)

type siteMapBuilder struct {
	domainURL   string
	visitedURLs set[string]
	urlsToVisit set[string]
}

func NewSiteMapBuilder(url string) *siteMapBuilder {
	return &siteMapBuilder{
		domainURL:   url,
		visitedURLs: set[string]{},
		urlsToVisit: set[string]{},
	}
}

func (builder *siteMapBuilder) Parse() ([]string, error) {
	builder.urlsToVisit.add(builder.domainURL)

	for len(builder.urlsToVisit) != 0 {
		url := builder.urlsToVisit.next()
		builder.urlsToVisit.remove(url)

		if builder.visitedURLs.has(url) {
			continue
		}

		err := builder.parseURL(url)
		if err != nil {
			return nil, err
		}
		builder.visitedURLs.add(url)
	}

	return setToSlice(builder.visitedURLs), nil
}

func (builder *siteMapBuilder) parseURL(url string) error {
	log.Printf("Client: GET %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	links, err := linkparser.Parse(resp.Body)
	if err != nil {
		return err
	}

	hrefs := getHRefs(links)

	for _, href := range hrefs {
		if !sameDomainLink(href, builder.domainURL) {
			continue
		}
		hrefWithDomain := formatHRef(href, builder.domainURL)
		if !builder.visitedURLs.has(hrefWithDomain) {
			builder.urlsToVisit.add(hrefWithDomain)
		}
	}

	return nil
}

func getHRefs(links []linkparser.Link) []string {
	hrefs := make([]string, len(links))
	for i, link := range links {
		hrefs[i] = link.Href
	}
	return hrefs
}

func formatHRef(url, domain string) string {
	ind := strings.Index(url, domain)
	if ind == -1 {
		return domain + url
	} else {
		return url
	}
}

func sameDomainLink(url, domain string) bool {
	if len(url) == 0 {
		return false
	}

	if url[0] == '/' {
		return true
	}

	if strings.Index(url, domain) == 0 {
		return true
	} else {
		return false
	}
}

type set[T comparable] map[T]struct{}

func (s set[T]) add(elem T) {
	s[elem] = struct{}{}
}

func (s set[T]) remove(elem T) {
	delete(s, elem)
}

func (s set[T]) has(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s set[T]) next() T {
	for elem := range s {
		return elem
	}
	var result T
	return result
}

func setToSlice(set set[string]) []string {
	slice := []string{}
	for key := range set {
		slice = append(slice, key)
	}
	return slice
}
