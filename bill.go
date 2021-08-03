package main

type bill struct {
	name string
	menuItems map[string]float64
	tip float64
}

func createBill (name string) bill {
	createdBill := bill{
		name: name,
		menuItems: map[string]float64{},
		tip: 0,
	}

	return createdBill
}