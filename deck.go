package main

import (
	"fmt"
	"strings"

	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck []string

func (d deck) write()  {
	for i, card := range d {
		fmt.Println(i ,card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return  d[:handSize], d[handSize:]
}

func (d deck) toString()string {
	return strings.Join([]string(d), ",\n")
}

func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename);
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1);
	}
	decksString := string(bs)
	deckStrings := strings.Split(decksString, ",\n");
	return deck(deckStrings);
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source);
	for i := range d {
		newPosition := r.Intn(len(d) -1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}