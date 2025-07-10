package main

import (
	"crawly/pkg"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <url> <thread count> <page limit>")
		log.Fatal("Invalid argument count")
	}

	maxWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("invalid MaxWorkers: %v", err)
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("invalid MaxPages: %v", err)
	}

	fmt.Printf("Starting crawl of: %s\n", os.Args[1])

	crawler := pkg.Crawler{
		Pages:   make(map[string]int),
		BaseURL: os.Args[1],
		Config: pkg.Config{
			MaxWorkers: maxWorkers,
			MaxPages:   maxPages,
		},
	}

	crawler.CrawlDomain()
	pkg.PrintOutput(crawler.Pages)
}
