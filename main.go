package main



//import "fmt"

func main() {
	//card := newCard()
	//fmt.Println(card)
	//cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	//cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	//cards = append(cards, "Six of Spades")
	//fmt.Println(cards)
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	//cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	//cards = append(cards, "Six of Spades")

	//cards := newDeck()
	//hand, remainingCards :=deal(cards, 5)
	//cards.write()
	//fmt.Println()
	//hand.write()
	//fmt.Println()
	//remainingCards.write()
	//cards.saveToFile("myCardsFile")
	deck := newDeckFromFile("myCardsFile")
	deck.shuffle()
	deck.write()

}

//func newCard() string {
//	return "Five of Diamonds"
//}
