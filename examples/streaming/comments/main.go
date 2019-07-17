package main

package main

import (
	"fmt"

	"github.com/mfdeux/snoo"
)

var client *snoo.Client

func main() {
	client = snoo.NewPublicClient("testClient 0.0.1")
	stream := client.StreamLinkComments("ce5s57", true)
	fmt.Println("About to start")
	stream.Start()
	fmt.Println("Started main")
	for link := range stream.Comments {
		fmt.Println(link)
	}
	fmt.Println("Stopped main")
	stream.Stop()
}
