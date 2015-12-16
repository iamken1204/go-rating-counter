package main

import (
	"fmt"
	web "github.com/iamken1204/rating-counter-web"
)

func main() {
	port := ":1234"
	fmt.Printf("Start server on: %s", port)
	web.Serve(":1234")
}
