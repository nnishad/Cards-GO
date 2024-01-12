package main

func main() {
	//cards := newDeck()

	//hand, remainingCards := deal(cards, 3)
	//hand.print()
	//remainingCards.print()

	//fmt.Print(cards.toString())
	//err := cards.saveToFile("my_deck")
	//if err != nil {
	//	return
	//}

	newDeck := newDeckFromFile("my_deck")
	//fmt.Println(newDeck)

	newDeck.shuffle()
	newDeck.print()
}
