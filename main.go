package main

import "fmt"

func main() {
	name:= "joey"
	myBill := createBill(name)
	fmt.Println(myBill.format())

}