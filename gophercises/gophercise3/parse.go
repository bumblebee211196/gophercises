package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href	string
	Text	string
}

func Parse(r io.Reader) ([]Link, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	linkNodes := getLinkNodes(node)
	for _, linkNode := range linkNodes {
		links = append(links, buildLink(getLinkNodeHref(linkNode), getLinkNodeText(linkNode)))
	}
	return links, nil
}

func buildLink(href, text string) Link {
	return Link{href, text}
}

func getLinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, getLinkNodes(c)...)
	}
	return nodes
}

func getLinkNodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	var ret string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret += getLinkNodeText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func getLinkNodeHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}
