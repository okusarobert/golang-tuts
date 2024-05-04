package main

import "fmt"

func updateName(x *string) {
	*x = "Okusa"
}

func main() {

	name := "Robert"

	// updateName(name)

	// fmt.Println("memory address of name is : ", &name)

	m := &name

	fmt.Println("memory address is: ", m)
	fmt.Println("value at memory address is: ", *m)

	// fmt.Println(name)

	updateName(m)

	fmt.Println(name)

}
