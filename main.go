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
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
