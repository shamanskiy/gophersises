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

	siteMap, err := sitemap.BuildMap(*url, reporter)
	if err != nil {
		log.Println(err)
	}
	printSiteMap(siteMap)
}

func createInterruptReporter() chan []string {
	interruptChannel := make(chan os.Signal, 1)
	reporter := make(chan []string)
	signal.Notify(interruptChannel, os.Interrupt)
	go func() {
		<-interruptChannel
		reporter <- []string{}
		siteMap := <-reporter
		printSiteMap(siteMap)
		os.Exit(1)
	}()

	return reporter
}

func printSiteMap(siteMap []string) {
	for _, url := range siteMap {
		fmt.Println(url)
	}
}
