package main

import "fmt"

// func main() {
// 	cards := newDeck()
// 	cards.shuffle()
// 	cards.print()
// }

type contact struct {
	address string
	phone   string
}

type person struct {
	firstName string
	lastName  string
	contact   contact
}

func main() {
	alex := person{firstName: "Archit", lastName: "Maheshwari", contact: contact{address: "Modinagar", phone: "243399"}}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	alexPointer := &alex
	alexPointer.updateName("Rajat")
	alex.print()

	// MAPS
	colors := map[string]string{
		"red":   "red",
		"green": "green",
		"white": "white",
	}

	// var colors map[string]string
	// colors := make(map[string]string)
	// fmt.Println(colors)
	// delete(colors, "green")
	printMap(colors)

}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
