// Add search depth restriction
// Add multi-threading
package sitemap

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/Shamanskiy/gophercises/base"
	"github.com/Shamanskiy/gophercises/linkparser"
)

type siteParser struct {
	domainURL   string
	visitedURLs base.Set[string]
	urlsToVisit base.Set[string]
}

// Parse domain and collect all reachable urls on the same domain
func ParseSite(domainUrl string, reporter chan []string) ([]string, error) {
	// Initialize the parser with domainUrl to visit
	parser := makeSiteParser(domainUrl)

	// An optional channel to get the intermediate result
	// while the site map parser is running
	if reporter != nil {
		launchReporter(&parser, reporter)
	}

	for !parser.urlsToVisit.Empty() {
		url := parser.urlsToVisit.Pop()

		if parser.visitedURLs.Has(url) {
			continue
		}
		parser.visitedURLs.Add(url)

		err := parser.parseURL(url)
		if err != nil {
			return parser.visitedURLs.ToSlice(), err
		}
	}

	return parser.visitedURLs.ToSlice(), nil
}

/* Returns the site map XML in this format:
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/foo1.html</loc>
  </url>
  <url>
    <loc>http://www.example.com/foo2.html</loc>
  </url>
</urlset>
*/
func MakeXmlMap(urls []string) (string, error) {
	siteMap := siteMapXmlFormat{Schema: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, url := range urls {
		siteMap.Urls = append(siteMap.Urls, urlXmlFormat{url})
	}

	siteMapXml := bytes.Buffer{}
	// add <?xml version="1.0" encoding="UTF-8"?>
	siteMapXml.Write([]byte(xml.Header))

	encoder := xml.NewEncoder(&siteMapXml)
	encoder.Indent("", "  ")

	err := encoder.Encode(siteMap)
	if err != nil {
		return "", err
	}

	return siteMapXml.String(), nil
}

type siteMapXmlFormat struct {
	XMLName xml.Name       `xml:"urlset"`
	Schema  string         `xml:"xmlns,attr"`
	Urls    []urlXmlFormat `xml:"url"`
}

type urlXmlFormat struct {
	Loc string `xml:"loc"`
}

func makeSiteParser(domainUrl string) siteParser {
	parser := siteParser{
		domainURL:   removeTrailingSlash(domainUrl),
		visitedURLs: base.Set[string]{},
		urlsToVisit: base.Set[string]{},
	}
	parser.urlsToVisit.Add(domainUrl)

	return parser
}

func launchReporter(parser *siteParser, reporter chan []string) {
	go func() {
		<-reporter
		reporter <- parser.visitedURLs.ToSlice()
	}()
}

func (parser *siteParser) parseURL(url string) error {
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
		if !sameDomainLink(href, parser.domainURL) {
			continue
		}
		hrefWithDomain := formatHRef(href, parser.domainURL)
		if !parser.visitedURLs.Has(hrefWithDomain) {
			parser.urlsToVisit.Add(hrefWithDomain)
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
