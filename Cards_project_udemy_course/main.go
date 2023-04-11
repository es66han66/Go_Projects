package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards1")

	// cards.print()

	cards.shuffle()

	// cards.print()

	cards.saveToFile("my_cards2")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()

	// remainingCards.print()

	// cards := getDeckFromFile("my_cards")

	// cards.print()
}
