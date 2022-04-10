package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(htmlDoc io.Reader) ([]Link, error) {
	doc, err := html.Parse(htmlDoc)
	if err != nil {
		return nil, err
	}

	return findLinks(doc), nil
}

func findLinks(node *html.Node) []Link {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []Link{buildLink(node)}
	}

	links := []Link{}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, findLinks(child)...)
	}
	return links
}

func buildLink(node *html.Node) Link {
	href := getHrefFromLink(node)
	text := getTextFromLink(node)
	return Link{href, strings.TrimSpace(text)}
}

func getHrefFromLink(node *html.Node) string {
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			return attribute.Val
		}
	}
	return ""
}

func getTextFromLink(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	linkText := ""
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		linkText += getTextFromLink(child)
	}
	return linkText
}
