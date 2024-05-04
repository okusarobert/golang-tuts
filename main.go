package main

import (
	"fmt"
	"math"
)

// void func
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// return values

func circleArea(r float64) float64 {
	return (math.Pi * r * r)
}
func main() {

	// sayGreeting("Robert")
	// cycleNames([]string{"Okusa", "Bonny", "Delice"}, sayGreeting)
	// cycleNames([]string{"Okusa", "Bonny", "Delice"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)

	// fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f \n", a1)
	fmt.Printf("circle 2 is %0.3f \n", a2)
}
