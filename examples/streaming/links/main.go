package main

import (
	"fmt"

	"github.com/mfdeux/snoo"
)

var client *snoo.Client

func main() {
	client = snoo.NewPublicClient("testClient 0.0.1")
	stream := client.StreamLinks("nba", true).Start()
	fmt.Println("Started main")
	for link := range stream.Links {
		fmt.Println(link)
	}
	fmt.Println("Stopped main")
	stream.Stop()
}
