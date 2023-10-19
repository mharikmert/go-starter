package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func helloWorldDeck() string {
	return "hello world from deck.go!"
}

// print function with parameters
func print(d deck) {
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
}

// returns a fresh deck
func newDeck() deck {
	cardSuits := []string{"Hearts", "Diamonds", "Spades"}
	cardValues := []string{"Ace", "Two", "Three"}

	cards := deck{} // empty slice

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardSuit+" of "+cardValue)
		}
	}

	return cards
}

// receiver function declaration
func (d deck) printReceiver() {
	for i, card := range d {
		fmt.Println(i, " ", card)
	}
}

// slice range syntax and multiple return values
func handle(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

// returns [] string to string with "," separator.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// string [] -> string -> [] byte
func (d deck) saveToFile(fileName string) {
	byteArrayOfDeck := []byte(d.toString())
	saveToFile(fileName, byteArrayOfDeck)
}

// [] byte -> string -> string []
func readDeckFromFile(fileName string) deck {
	byteArrayOfDeck, _ := readFromFile(fileName)

	// string array of the byte slice
	s := strings.Split(string(byteArrayOfDeck), ",")

	// return s as a deck
	return deck(s)
}

func (d deck) shuffleWithStandardMathLibrary() {
	// generate a random number (card position) from zero to length of slice
	currentPosition := rand.Intn(len(d) - 1)

	// shuffle the deck by replacing the card in the current position(index) with the card in the remaining index from the length of the slice
	for i := range d {
		// swap operation by simply using comma
		d[i], d[currentPosition] = d[currentPosition], d[i]
	}
}

func (d deck) shuffle() {
	// returns an int64 number as the random seed resource
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	currentPosition := r.Intn(len(d) - 1)
	for i := range d {
		// swap operation by simply using comma
		d[i], d[currentPosition] = d[currentPosition], d[i]
	}
}
