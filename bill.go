package main

import "fmt"

type Bill struct {
	name string
	menuItems map[string]float64
	tip float64
}

func newBill (name string) Bill {
	createdBill := Bill{
		name: name,
		menuItems: map[string]float64{},
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

	formatedString += fmt.Sprintf("%-25v \n","--------------------")

	formatedString += fmt.Sprintf("%-25v ... $%v \n","tips:", bill.tip)
	formatedString += fmt.Sprintf("%-25v ... $%0.2f", "total price::", totalPrice + bill.tip)
	return formatedString
}

func (bill *Bill) updatingTips (updatedBillNumber float64) {
	memoryOfTips := &bill.tip
	*memoryOfTips = updatedBillNumber
	fmt.Printf("------------------------------------------heeeyyy:::", updatedBillNumber)
}

func (bill *Bill) addMenuItem (name string, price float64) {
	bill.menuItems[name] = price
}