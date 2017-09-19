package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}
	if d[15] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %v", d[15])
	}

}
func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	ld := newDeckFromFile("_decktesting")

	if len(ld) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(ld))
	}

	if ld[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", ld[0])
	}
	if ld[15] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %v", ld[15])
	}
	os.Remove("_decktesting")
}
