package main

import (
	"fmt"
	"log"

	"github.com/mfdeux/snoo"
)

var client *snoo.Client

func main() {
	client = snoo.NewPublicClient("testClient 0.0.1")
	links, err := client.GetNewLinks("nba")
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		fmt.Println(link.Title, link.CreatedUTC)
	}
}
