package main

import (
	"fmt"
	"sort"
)

func main() {
	// greetings := "hello there friends"
	// fmt.Println(strings.Contains(greetings, "hellookusa"))

	// the original string remains unchanged
	// fmt.Println(strings.ReplaceAll(greetings, "hello", "hi"))

	// fmt.Println(strings.Split(greetings, " "))

	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "th"))

	ages := []int{23, 45, 67, 89, 90, 34, 56}
	sort.Ints(ages) // changes the original slice
	// fmt.Println(ages)

	names := []string{"yoshi", "robert", "tommy", "solo", "eddie"}

	// index := sort.SearchInts(ages, 100)

	// fmt.Println(index)

	sort.Strings(names)

	// fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "robert"))
}
