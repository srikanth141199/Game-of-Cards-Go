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
