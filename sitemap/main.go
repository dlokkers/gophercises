package main

import (
	"encoding/xml"
	"fmt"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dlokkers/gophercises/link/parser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Loc string `xml:"loc"`
}

type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []loc  `xml:"url"`
}

func main() {
	htmlPagePtr := flag.String("page", "https://gophercises.com", "html page to parse")
	depthPtr    := flag.Int("depth", 3, "amount of clicks to traverse the site")
	flag.Parse()

	links := bfs( *htmlPagePtr, *depthPtr )

	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, l := range links {
		toXml.Urls = append(toXml.Urls, loc{l})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "    ")
	if err := enc.Encode(toXml); err != nil {
		log.Fatal(err)
	}
}

func bfs( u string, d int ) []string {
	seen := make(map[string]struct{})
	q    := []string{}
	nq   := []string{u}

	for i := 0; i <= d; i++ {
		q, nq = nq, []string{}
		for _, url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			nq = getHrefs(url)
		}
	}

	var ret []string
	for k, _ := range seen {
		ret = append(ret, k)
	}
	return ret
}

func getHrefs( u string ) []string { 
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	n, err := parse.FirstNode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	l := parse.ParseHTML( n )

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	var links []string
	for _, h := range l {
		links = append(links, h.Href)
	}

	return filterHrefs( links, baseUrl.String() )
}

func filterHrefs( links []string, base string ) []string {
	var ret []string
	for _, l := range links {
		switch{
			case strings.HasPrefix(l, "/"):
				ret = append(ret, base + l)
			case strings.HasPrefix(l, "."):
				ret = append(ret, base + "/" + l)
			case strings.HasPrefix(l, base):
				ret = append(ret, l)
		}
	}
	return unique(ret)
}

func unique( s []string ) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append( list, entry)
		}
	}
	return list
}
