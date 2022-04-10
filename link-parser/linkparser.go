package linkparser

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlDoc []byte) (map[string]string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlDoc))
	if err != nil {
		return nil, err
	}

	links := map[string]string{}
	parseHtml(doc, links)

	return links, nil
}

func parseHtml(node *html.Node, links map[string]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		link := node.Attr[0].Val
		linkText := parseLink(node)
		links[link] = strings.TrimSpace(linkText)
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		parseHtml(child, links)
	}
}

func parseLink(node *html.Node) string {
	linkText := ""

	if node.Type == html.TextNode {
		linkText += node.Data
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		linkText += parseLink(child)
	}

	return linkText
}
