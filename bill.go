package main

import "fmt"

type Bill struct {
	name string
	menuItems map[string]float64
	tip float64
}

func createBill (name string) Bill {
	createdBill := Bill{
		name: name,
		menuItems: map[string]float64{
		"pizza": 12.76,
		"burger": 10.99,
		"shawarma": 3,
		"cola": 1.99,
	},
		tip: 0,
	}

	return createdBill
}

func (bill Bill) format () string{
	formatedString:= "Bill breakdown \n"
	var totalPrice float64 = 0
	for key,value := range bill.menuItems{
		formatedString += fmt.Sprintf("%-25v ... $%v \n", key+":", value)
		totalPrice += value
	}

	formatedString += fmt.Sprintf("%-25v ... $%0.2f", "total price::", totalPrice)
	return formatedString
}