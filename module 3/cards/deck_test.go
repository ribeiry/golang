package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected  deck length of 20, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card  of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubes" {
		t.Errorf("Expected last card of Four of Clubes, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()
	d.saveToFile("_decktesting.txt")

	loadDeck := newDeckFromFile("_decktesting.txt")

	if len(loadDeck) != 16 {
		t.Errorf("Expected  deck length of 20, but got %d", len(loadDeck))
	}
	os.Remove("_decktesting.txt")
}
