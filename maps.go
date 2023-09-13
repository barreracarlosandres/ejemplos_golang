package main

import "fmt"

func mapExample1() {
	// make(map[key-type]val-type).
	colors := map[string]string{
		"red":   "code-red",
		"green": "code-green",
	}
	fmt.Println(colors)
	printMap(colors)
}

func mapExample2() {
	var colors map[string]string
	fmt.Println(colors)
}

func mapExample3() {
	colors := make(map[string]string)
	fmt.Println(colors)
}

func mapExample4() {
	colors := make(map[string]string)
	colors["white"] = "code-white"
	fmt.Println(colors)
}

func mapExample5() {
	colors := make(map[string]string)
	colors["white"] = "code-white"
	fmt.Println(colors)
	delete(colors, "white")
	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex	:= range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}