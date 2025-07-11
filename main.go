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

	baseURL, err := pkg.NormalizeURL(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	schemaURL := pkg.ValidateScheme(baseURL)

	fmt.Printf("Starting crawl of: %s\n", schemaURL)

	crawler := pkg.Crawler{
		Pages:   make(map[string]int),
		BaseURL: schemaURL,
		Config: pkg.Config{
			MaxWorkers: maxWorkers,
			MaxPages:   maxPages,
		},
	}

	crawler.CrawlDomain()
	pkg.PrintOutput(crawler.Pages)
}
