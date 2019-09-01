package parse

import (
	"io"
//	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func FirstNode( r io.Reader ) (*html.Node, error) {
	return html.Parse(r)
}

func ParseHTML(n *html.Node) ([]Link) {
	var links []Link

	if n.Type == html.ElementNode && n.Data == "a" {
		return append(links, parseLinkNode(n))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append( links, ParseHTML(c)...)
	}


	return links
}

func parseLinkNode( n *html.Node ) Link {
	var l Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			l.Href =  a.Val
		} 
	}

	b := retrieveText(n)
	l.Text = strings.TrimSpace(b)

	return l
}

func retrieveText(n *html.Node) (string) {
	var s string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode && c.Type != html.CommentNode {
			s = s + c.Data
		}
		s = s + retrieveText(c)
	}
	return s
}
