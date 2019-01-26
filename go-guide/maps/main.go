package main

import "fmt"

func main(){
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	// var colors map[string]string

	//Create a map and delete an element inside of it

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	// delete(colors, 10)
	
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}