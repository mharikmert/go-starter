package main

import (
	"fmt"
)

// variable declaration
var card string

func helloWorld() string {
	return "Hello World"
}

func main() {

	fmt.Println(helloWorld())
	fmt.Println(helloWorldDeck())

	card = "Previously declared string"

	// variable declaration and assignment
	newlyDeclared := "Newly declared string"

	fmt.Println(card + "\n" + newlyDeclared)

	// direct access to the helper.go file, since it's in the same package(main)
	fmt.Println(sum(3, 5))

	// SLICES: Slices are like arrays, but they can resized.
	cards := []string{"some", "other"}
	messages := []string{helloWorld(), "Good Bye World"}

	fmt.Println(cards)
	fmt.Println(messages)

	// FOR LOOP
	for i, message := range messages {
		fmt.Println(i, message)
	}

	// slice append function
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	newDeck := newDeck()

	// receiver usage
	fmt.Println("\n-----RECEIVER FUNCTION OUTPUT-----")
	newDeck.printReceiver()

	fmt.Println("\n-----STANDARD FUNCTION OUTPUT-----")
	// standard function call
	print(newDeck)

	// BYTE SLICES: Computer friendly way of thinking about a string
	byteSlice := []byte("byteSlice")
	fmt.Println(byteSlice)

	fmt.Println(newDeck.toString())

	fileName := "my_cards.txt"
	newDeck.saveToFile(fileName)

	currentDeck := readDeckFromFile(fileName)
	fmt.Println("\n-----DECK FROM THE FILE-----")
	// standard function call
	currentDeck.printReceiver()

	fmt.Println("\n-----SHUFFLED DECK WITH STDLIB-----")
	currentDeck.shuffleWithStandardMathLibrary()
	currentDeck.printReceiver()

	fmt.Println("\n-----SHUFFLED DECK WITH CUSTOM RANDOM SOURCE-----")
	currentDeck.shuffle()
	currentDeck.printReceiver()

}
