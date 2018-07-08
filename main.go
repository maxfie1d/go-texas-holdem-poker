package main

import (
	"fmt"
)



func main() {
	var cards []Card

	kinds := []CardKind{Diamond, Heart, Clover, Spade}
	for _, kind := range kinds {
		for _, v := range []int {1,2,3,4,5,6,7,8,9,10,11,12,13} {
			card := Card{kind: kind, number:v}
			cards = append(cards, card)
		}
	}

	for _, c := range cards {
		fmt.Println(c)
	}
}
