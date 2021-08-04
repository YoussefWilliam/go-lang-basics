package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readingInputs(promt string, reader* bufio.Reader)(string, error){
	fmt.Print(promt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promtOptions(b Bill){
	inputReader := bufio.NewReader(os.Stdin)
	options, _ :=readingInputs("Choose an option ('a' -> add item, 's' -> save bill, 't' -> add tips): \n",inputReader)
	fmt.Println(options)
}

func createBill() Bill {
	inputReader := bufio.NewReader(os.Stdin)

	name, _ :=readingInputs("Create a new bill, what's your name? \n",inputReader)

	bill := newBill(name)

	fmt.Printf("Created the bill under %q name \n", name)

	promtOptions(bill)
	return bill
}
func main() {

	myBill := createBill()
	fmt.Println("my new bill",myBill)
	// name:= "joey"
	// myBill := createBill(name)
	// fmt.Println(myBill.format())

	// fmt.Println("Original tips::", myBill.tip)
	// myBill.updatingTips(4)
	// fmt.Println("Updated tips::", myBill.tip)

	// myBill.addMenuItem("beer", 20)
	// fmt.Println(myBill.format())

}