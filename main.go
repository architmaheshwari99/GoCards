package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	// Slice
	cards := []string{"Archit", "Casper", newCard()}

	// It doesn't modify the existing slice, returns a new slice
	cards = append(cards, "Anupam")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
