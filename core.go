package main

import (
	rc "github.com/iamken1204/rating-counter"
)

func main() {
	targets := rc.GetTargets()
	for _, target := range targets {
		rc.Crawl(target.ID, "google", target.Keyword, target.Url)
		rc.Crawl(target.ID, "yahoo", target.Keyword, target.Url)
	}
}
