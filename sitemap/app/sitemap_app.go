package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shamanskiy/gophercises/sitemap"
)

func main() {
	url := flag.String("url", "http://127.0.0.1", "domain url of the site that we want to build a site map for")
	flag.Parse()

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)
	go func() {
		<-interruptChannel
		fmt.Println("Interrupted")
		os.Exit(1)
	}()

	siteMap, err := sitemap.BuildMap(*url)
	if err != nil {
		log.Fatalln(err)
	}

	for _, url := range siteMap {
		fmt.Println(url)
	}
}
