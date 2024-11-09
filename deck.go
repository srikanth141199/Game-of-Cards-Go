package main

import "fmt"

type deck []string

//below is an example of receiver function
//by declaring like this we can say any variable inside the application with type deck can access the function print

func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
