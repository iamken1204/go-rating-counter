package main

import (
	"fmt"
	web "github.com/iamken1204/rating-counter-web"
)

func main() {
	port := ":8080"
	fmt.Printf("Start server on: %s", port)
	web.Serve(port)
}
