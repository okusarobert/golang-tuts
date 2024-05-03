package main

import "fmt"

func main() {
	// arrays
	// var ages [3]int = [3]int{20, 25, 35}
	var ages = [3]int{20, 25, 35}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "rachael"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)

	var scores = []int{100, 50, 60}

	scores[2] = 90
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]

	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")

	fmt.Println(rangeOne)

}
