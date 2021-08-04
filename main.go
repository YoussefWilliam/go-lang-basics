package main

import "fmt"

func main() {
	name:= "joey"
	myBill := createBill(name)
	// fmt.Println(myBill.format())

	fmt.Println("Original tips::", myBill.tip)
	myBill.updatingTips(4)
	fmt.Println("Updated tips::", myBill.tip)

	myBill.addMenuItem("beer", 20)
	fmt.Println(myBill.format())

}