package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spaces" //:= only during initialization
	// card = "Five of Diamonds"
	// card := newCard()

	// fmt.Println(card)
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spaces")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// cards.print()
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// fmt.Println(cards.toString())
	cards := newDeck()
	cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
