package maps

import "fmt"

// create a empty map
var cities = make(map[string]int)

func createMap() {
	fmt.Println("* Execute func createMap")

	// createa a map with values
	colors := map[string]string{
		"red":   "code-red",
		"green": "code-green",
	}
	fmt.Println(colors)

	fmt.Println()
}

func addValues() {
	fmt.Println("* Execute func addValues")

	fmt.Println("before add values in 'cities':", cities)
	cities["cali"]   = 1
	cities["bogotá"] = 2
	cities["buga"]   = 3
	fmt.Println("All values in 'cities':", cities)
	fmt.Println("Value of key 'buga':", cities["buga"] )

	fmt.Println()
}

func createEmptyMap() {
	fmt.Println("* Execute func createEmptyMap")
	colors := make(map[string]string)
	fmt.Println(colors)

	fmt.Println()
}

func deleteValue() {
	fmt.Println("* Execute func example5")

	fmt.Println("before delete :",cities)
	delete(cities, "buga")
	fmt.Println("after delelte cities:", cities)

	fmt.Println()
}

func printMap(c map[string]string) {
	fmt.Println("* Execute func printMap")
	for color, hex	:= range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func Example() {
	fmt.Println("------------ Maps Examples ------------")
	createMap()
	addValues()
	createEmptyMap()
	deleteValue()
}