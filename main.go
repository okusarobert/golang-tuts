package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() bill {

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill: ", reader)
	b := newBill(name)
	fmt.Print("Created bill - ", b.name, "\n")

	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, s - save the bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		convertedPrice, error := strconv.ParseFloat(strings.TrimSpace(price), 64)
		if error != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, convertedPrice)

		fmt.Println("Item added ", name, price)
		promptOptions(b)
		break
	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
		break
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		convertedTip, error := strconv.ParseFloat(strings.TrimSpace(tip), 64)
		if error != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(convertedTip)
		promptOptions(b)
		break
	default:
		fmt.Println("Invalid option ...")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()

	promptOptions(&myBill)

	// myBill.format()

	// myBill.updateTip(10)

	// myBill.addItem("pie", 5.99)
	// myBill.addItem("cake", 2.99)
	// myBill.addItem("soda", 1.99)
	// myBill.addItem("ice cream", 0.51)

	fmt.Println(myBill.format())
}
