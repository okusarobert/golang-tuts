package main

import "fmt"

func main() {
	// menu := map[string]float64{
	// 	"soup":           4.99,
	// 	"pie":            7.99,
	// 	"salad":          6.99,
	// 	"toffee padding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["soup"])

	// looping maps
	// for key, value := range menu {
	// 	fmt.Println(key, "-", value)
	// }

	// ints as key type

	phonebook := map[int]string{
		256759847505: "Robert",
		256759847501: "Robert 1",
		256759847502: "Robert 2",
		256759847503: "Robert 3",
	}

	// for key, value := range phonebook {
	// 	fmt.Println(key, "-", value)
	// }

	phonebook[256759847503] = "bowser"

	fmt.Println(phonebook)

}
