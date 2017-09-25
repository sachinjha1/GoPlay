package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("ping")
		c <- "ping----->"
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("pong")
		c <- "<-----pong"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	/*
		With below line commented Ping/Pong works fine alternate ping pong in console. => Ping Pong Ping Pong
		Once i uncomment below line Ping Pong order breaks... => Ping Ping Pong Ping

		Why ??????? My understanding is here I dont need multiple CPU so goMaxProc 1 should still work fine
	*/
	//runtime.GOMAXPROCS(1)
	c := make(chan string, 10)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
