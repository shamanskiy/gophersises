package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Shamanskiy/gophercises/sitemap"
)

func main() {
	url := flag.String("url", "http://127.0.0.1", "domain url of the site that we want to build a site map for")
	flag.Parse()

	builder := sitemap.NewSiteMapBuilder(*url)
	siteMap, err := builder.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	for _, url := range siteMap {
		fmt.Println(url)
	}
}
