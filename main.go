package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// var temp = 123
	// tempStr := "this is a String"

	// fmt.Println(temp)
	// fmt.Println(tempStr)

	//cards := []string{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	//fmt.Println(cards)

	cards = append(cards, "six of Spades")
	fmt.Println((cards))

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Ace of Spades"
}
