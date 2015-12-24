package main

import (
	"fmt"
	"os"

	rc "github.com/iamken1204/rating-counter"
)

func main() {
	targets := rc.GetTargets()
	if len(targets) == 0 {
		fmt.Println("No targets to be crawled.")
		os.Exit(1)
	}
	for _, target := range targets {
		rc.Crawl(target.ID, "google", target.Keyword, target.Url)
		rc.Crawl(target.ID, "yahoo", target.Keyword, target.Url)
	}
}
