package main

import "fmt"

func updateName(x string) {
	x = "Robert"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 1.99
}

func main() {

	// group A types --> strings, ints, bools, floats, arrays, structs

	// group B types --> slice, maps, functions

	// name := "Okusa"

	// updateName(name)

	// fmt.Println(name)

	menu := map[string]float64{
		"ice cream": 3.55,
		"pie":       5.95,
	}

	updateMenu(menu)

	fmt.Println(menu)

}
