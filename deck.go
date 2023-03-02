package main

import "fmt"

// Create a new type of dec
// which is a slice of strings

type deck []string

// Receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
