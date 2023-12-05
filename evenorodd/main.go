package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := [10]int{}

	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(10)
		list[i] = randomNum
	}

	for _, value := range list {
		if value % 2 == 0 {
			fmt.Printf("%v is even\n", value)
		} else {
			fmt.Printf("%v is odd\n", value)
		}
	}
}