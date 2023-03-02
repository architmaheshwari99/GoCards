package main

func main() {
	cards := newDeck()
	cards.saveToFile("mycards")
	deck := newDeckFromFile("mycards")
	deck.print()
}
