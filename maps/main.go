package main

import "fmt"

func main() {
	// First way of declaring map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Second way of declaring map
	// var colors map[string]string

	// Third way of declaring map
	// colors := make(map[string]string)

	// Always use square brace syntax with maps, do not use . like other languages
	// colors["white"] = "#ffffff"

	// Delete a key/value from map
	// delete(colors, "white")

	// iterate over key/value pairs of map
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
