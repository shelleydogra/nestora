package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Welcome to Nestora

type Property struct {
	Address    string
	UnitNumber string
	TenantName string
	RentAmount float64
	Paid       bool
}

var properties []Property

func main() {
	for {
		fmt.Println("=============================")
		fmt.Println("      Welcome to Nestora      ")
		fmt.Println("=============================")
		fmt.Println("1. Add Property")
		fmt.Println("2. List Properties")
		fmt.Println("3. Mark Rent as Paid")
		fmt.Println("4. Show Rent Roll Report")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addProperty()
		case 2:
			listProperty()
		case 3:
			markRentPaid()
		case 4:
			showRentRoll()
		case 5:
			fmt.Print("Exiting Nestora. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addProperty() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Add New Property ---")

	fmt.Print("Enter property address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	fmt.Print("Enter unit number: ")
	unit, _ := reader.ReadString('\n')
	unit = strings.TrimSpace(unit)

	fmt.Print("Enter tenant name: ")
	tenant, _ := reader.ReadString('\n')
	tenant = strings.TrimSpace(tenant)

	fmt.Print("Enter monthly rent amount: ")
	var rent float64
	fmt.Scanln(&rent)

	newProperty := Property{
		Address:    address,
		UnitNumber: unit,
		TenantName: tenant,
		RentAmount: rent,
		Paid:       false,
	}

	properties = append(properties, newProperty)
	fmt.Println("Property added successfully!\n")
}

func listProperty() {
	fmt.Println("\n--- List of Properties ---")

	if len(properties) == 0 {
		fmt.Println("No properties found.")
		return
	}

	for i, p := range properties {
		fmt.Printf("%d) Address: %s | Unit: %s | Tenant: %s | Rent: %.2f | Paid: %v\n",
			i+1, p.Address, p.UnitNumber, p.TenantName, p.RentAmount, p.Paid)
	}
	fmt.Println()
}

func markRentPaid() {
	fmt.Println("\n--- Mark Rent as Paid ---")

	if len(properties) == 0 {
		fmt.Println("No properties to update.")
		return
	}

	for i, p := range properties {
		fmt.Printf("%d) Address: %s | Unit: %s | Tenant: %s | Paid: %v]n",
			i+1, p.Address, p.UnitNumber, p.TenantName, p.Paid)
	}

	var choice int
	fmt.Print("Enter the number of the property to mark as paid: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(properties) {
		fmt.Println("Invalid selection.")
		return
	}

	properties[choice-1].Paid = true
	fmt.Println("Rent marked as Paid!\n")
}

func showRentRoll() {
	fmt.Println("\n--- Rent Roll Report ---")

	if len(properties) == 0 {
		fmt.Println("No properties to show.")
		return
	}

	var totalRent float64
	var totalCollected float64

	for _, p := range properties {
		totalRent += p.RentAmount
		if p.Paid {
			totalCollected += p.RentAmount
		}

		status := "Unpaid"
		if p.Paid {
			status = "Paid"
		}

		fmt.Printf("Tenant: %s| Address: %s Unit %s | Rent: $%.2f | Status: %s\n",
			p.TenantName, p.Address, p.UnitNumber, p.RentAmount, status)
	}

	fmt.Println("\nSummary:")
	fmt.Printf("Total Rent Expected: $%.2f\n", totalRent)
	fmt.Printf("Total Rent Collectied: $%.2f\n", totalCollected)
	fmt.Printf("Outstanding Balance: $%.2f\n", totalRent-totalCollected)
	fmt.Println()
}
