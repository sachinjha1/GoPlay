package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(data[:count]))
}
