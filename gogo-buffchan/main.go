package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	names := make(chan string, 1)

	go printStar(names)
	go printHip(names)
	time.Sleep(2000 * time.Millisecond)
}

func printStar(names chan string) {
	for i := 0; i < 100; i++ {
		fmt.Print(i, "*")
	}
}

func printHip(names chan string) {
	for i := 0; i < 100; i++ {
		fmt.Print(i, "-")
	}
}
