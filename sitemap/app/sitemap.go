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

	reporter := createInterruptReporter()

	urls, err := sitemap.ParseSite(*url, reporter)
	if err != nil {
		log.Println(err)
	}

	printURLs(urls)
}

func createInterruptReporter() chan []string {
	interrupter := make(chan os.Signal, 1)
	signal.Notify(interrupter, os.Interrupt)

	reporter := make(chan []string)
	go func() {
		<-interrupter
		reporter <- []string{}
		collectedUrlsSoFar := <-reporter
		printURLs(collectedUrlsSoFar)
		os.Exit(1)
	}()

	return reporter
}

func printURLs(urls []string) {
	for _, url := range urls {
		fmt.Println(url)
	}
	fmt.Printf("Found %d URLs\n", len(urls))
}
