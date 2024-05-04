package main

import "fmt"

func main() {

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less then 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else if age >= 45 {
	// 	fmt.Println("age is greater than or equal to 45")
	// }

	names := []string{"yoshi", "okusa", "robert", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Printf("breaking at pos %v of value %v \n", index, value)
			break
		}
		fmt.Printf("the value at %v is %v \n", index, value)
	}

}
