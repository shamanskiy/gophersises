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
	interrupter := make(chan os.Signal, 1)
	signal.Notify(interrupter, os.Interrupt)

	reporter := make(chan []string)
	go func() {
		<-interrupter
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
