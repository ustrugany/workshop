package main

import (
	"fmt"
	"time"
)

func miner(name string, channel <-chan string) {
	for crypto := range channel {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Printf("%s mined %q\n", name, crypto)
	}
}

func broker(channel chan<- string, crypto []string) {
	for _, c := range crypto {
		fmt.Printf("sending %q\n", c)
		channel <- c
	}
	close(channel)
}

func main() {
	crypto := []string{
		"bitcoin",
		"litecoin",
		"ripple",
		"primecoin",
		"titcoin",
		"verge",
		"stellar",
		"gridcoin",
	}
	channel := make(chan string)
	// starting lightweight thread
	go miner("piotr", channel)
	go miner("simon", channel)
	// receiver needs to start before sender
	go broker(channel, crypto)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("all mined")
}
