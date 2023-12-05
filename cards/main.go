package main

import "fmt"

func main() {
	cards, err := readFile("my_cards")
	if err != nil {
		fmt.Println(err)
	}
	cards.print()

}