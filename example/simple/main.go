package main

import (
	"fmt"

	"github.com/leonhfr/anki-connect-go"
)

func main() {
	myClient := anki.NewDefaultClient()
	deckNames, err := myClient.DeckNames()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the deck names
	fmt.Println("Deck Names:", deckNames)
}
