package main

import (
	"fmt"
)

func main() {

	// other ways of declaring variables
	//var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{

		"red":   "#FF0000",
		"black": "#000000",
		"blue":  "#DER000",
	}

	// add a new entry
	colors["white"] = "#ffffff"

	// delete an entry from the map
	delete(colors, "white")

	printMap(colors)
}

// iterating over map objects
func printMap(c map[string]string) {

	for color, hex := range c {
		fmt.Println("Hex code of", color, "is", hex)
	}
}
