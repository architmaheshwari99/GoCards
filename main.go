package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	// Slice
	cards := deck{"Archit", "Casper", newCard()}

	// It doesn't modify the existing slice, returns a new slice
	cards = append(cards, "Anupam")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
