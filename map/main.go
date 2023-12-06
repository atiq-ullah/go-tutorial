package main

import "fmt"

func main() {
	// colors := make(map[string]string) 
	// var colors map[string]string

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#74500",
		"blue": "#dfdfefd",
	}

	colors["grey"] = "#342300"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}