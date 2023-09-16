package maps

import "fmt"

func Example1() {
	// make(map[key-type]val-type).
	colors := map[string]string{
		"red":   "code-red",
		"green": "code-green",
	}
	fmt.Println(colors)
	printMap(colors)
}

func Example2() {
	var colors map[string]string
	fmt.Println(colors)
}

func Example3() {
	colors := make(map[string]string)
	fmt.Println(colors)
}

func Example4() {
	colors := make(map[string]string)
	colors["white"] = "code-white"
	fmt.Println(colors)
}

func Example5() {
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