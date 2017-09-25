package main

import (
	"fmt"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(2)
	c := make(chan string)
	go ping(c)
	go pong(c)
	c <- "ping"
	time.Sleep(10000 * time.Millisecond)
}

func ping(c chan string) {
	for {
		pong := <-c
		fmt.Println(pong)
		time.Sleep(1000 * time.Millisecond)
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		ping := <-c
		fmt.Println(ping)
		time.Sleep(1000 * time.Millisecond)
		c <- "pong"
	}
}
