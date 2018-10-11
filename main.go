package main

import (
	"fmt"
)

func main() {
	cards := newDeckFromFile("new_deck")
	fmt.Println(cards)
}
