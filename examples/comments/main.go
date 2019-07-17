package main

import (
	"fmt"
	"log"

	"github.com/mfdeux/snoo"
)

var client *snoo.Client

func main() {
	client = snoo.NewPublicClient("testClient 0.0.1")
	links, err := client.GetLinkComments("cdyidp")
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		fmt.Println(link.Body)
	}
}
