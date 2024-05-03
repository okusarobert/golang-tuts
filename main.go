package main

import "fmt"

func main() {

	age := 33

	name := "robert"

	// Print
	fmt.Print("hello, ")
	fmt.Print("World! \n")

	// Printf() Formatted strings %_ format specifier

	fmt.Println("my age is ", age, " and my name is ", name)

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)

	fmt.Printf("age is of type %T \n", age)

	fmt.Printf("you scored %0.1f points! \n", 222.55)

	// Sprintf saving variable to string
	str := fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	fmt.Println("saved string is:  ", str)

}
