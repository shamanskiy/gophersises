// Add search depth restriction
// Add XML output
// Add multi-threading
// Add interruption output
package sitemap

import (
	"net/http"
	"strings"

	"github.com/Shamanskiy/gophercises/base"
	"github.com/Shamanskiy/gophercises/linkparser"
)

type siteMapBuilder struct {
	domainURL   string
	visitedURLs base.Set[string]
	urlsToVisit base.Set[string]
}

func BuildMap(domainUrl string) ([]string, error) {
	domainUrl = removeTrailingSlash(domainUrl)

	builder := siteMapBuilder{
		domainURL:   domainUrl,
		visitedURLs: base.Set[string]{},
		urlsToVisit: base.Set[string]{},
	}
	builder.urlsToVisit.Add(domainUrl)

	for !builder.urlsToVisit.Empty() {
		url := builder.urlsToVisit.Pop()

		if builder.visitedURLs.Has(url) {
			continue
		}
		builder.visitedURLs.Add(url)

		err := builder.parseURL(url)
		if err != nil {
			return builder.visitedURLs.ToSlice(), err
		}
	}

	return builder.visitedURLs.ToSlice(), nil
}

func (builder *siteMapBuilder) parseURL(url string) error {
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
		if !builder.visitedURLs.Has(hrefWithDomain) {
			builder.urlsToVisit.Add(hrefWithDomain)
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
		url = domain + url
	}

	return removeTrailingSlash(url)
}

func removeTrailingSlash(url string) string {
	if len(url) == 0 {
		return url
	}

	if url[len(url)-1:] == "/" {
		return url[:len(url)-1]
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
