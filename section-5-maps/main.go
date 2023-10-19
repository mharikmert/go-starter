package main

import "fmt"

func main() {
	var colors map[string]string // empty map declaration

	colors2 := make(map[string]string) // another way of declaring an empty map.

	colors3 := map[string]string{ // declaring with values
		"red":   "#234234",
		"green": "#12039",
		"blue":  "#23948",
	}
	fmt.Println(colors, colors2, colors3)

	// deleting a key value pair from the map:
	delete(colors3, "red")

	fmt.Println("printed value inside the colors3 map:", colors3["blue"]) // calling the key value pair by -> mapName["${key}""]

	// colors["red"] = "#090909"
	colors3["red"] = "something different "

	// fmt.Println(colors3)

	// fmt.Println(colors3)
	print(colors3)

	colors2["red"] = "give value to red" // you can assign the new value for this declaration. make()
	// change(colors2)
	print(colors2)

	// colors["red"] = "brand new" // you cannot assign a new key value pair for this declaration. var
	// print(colors)

}

func print(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func change(m map[string]string) {
	m["red"] = "red color"
}
