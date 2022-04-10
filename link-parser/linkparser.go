package linkparser

import (
	"bytes"

	"golang.org/x/net/html"
)

func ExtractLinks(htmlDoc []byte) (map[string]string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlDoc))
	if err != nil {
		return nil, err
	}

	links := map[string]string{}
	parsingFunc(doc, links)

	return links, nil
}

func parsingFunc(node *html.Node, links map[string]string) {
	if node.Type == html.ElementNode && node.Data == "a" {
		link := node.Attr[0].Val
		linkText := node.FirstChild.Data
		links[link] = linkText
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		parsingFunc(child, links)
	}
}
