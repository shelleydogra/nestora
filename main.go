package main

import (
	"fmt"
)

// Welcome to Nestora

type Property struct {
	Address    string
	UnitNumber string
	TenantName string
	RentAmount float64
	Paid       bool
}

func main() {
	fmt.Println("=============================")
	fmt.Println("      Welcome to Nestora      ")
	fmt.Println("=============================")
	fmt.Println("1. Add Property")
	fmt.Println("2. List Properties")
	fmt.Println("3. Mark Rent as Paid")
	fmt.Println("4. Show Rent Roll Report")
	fmt.Println("5. Exit")
}
