package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	delete(colors, "red")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hexa := range c {
		fmt.Println(color, " means ", hexa)
	}
}
