package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// --- Structs ---

type Property struct {
	ID      string
	Name    string
	Address string
	Units   []Unit
}

type Unit struct {
	ID         string
	UnitNumber string
	Bedrooms   int
	Bathrooms  float64
	SquareFeet int
	Leases     []Lease
}

type Tenant struct {
	ID       string
	FullName string
	Email    string
	Phone    string
}

type Lease struct {
	ID              string
	Tenant          Tenant
	StartDate       time.Time
	EndDate         time.Time
	MonthlyRent     float64
	SecurityDeposit float64
	RentHistory     []RentPayment
	Status          string // active, upcoming, ended
}

type RentPayment struct {
	Month      time.Time
	AmountDue  float64
	AmountPaid float64
	PaidDate   *time.Time
	Status     string // paid, unpaid, partial
	Notes      string
}

// --- Global Variables ---

var properties []Property

const dataFile = "data.json"

// --- Main Program ---

func main() {
	loadData()

	for {
		fmt.Println("=============================")
		fmt.Println("      Welcome to Nestora      ")
		fmt.Println("=============================")
		fmt.Println("1. Add Property")
		fmt.Println("2. List Properties")
		fmt.Println("3. Add Unit to Property")
		fmt.Println("4. List Units for a Property")
		fmt.Println("5. Create Lease for Unit")
		fmt.Println("6. List Leases for Unit")
		fmt.Println("7. Record Rent Payment")
		fmt.Println("8. Generate Rent Roll Report")
		fmt.Println("9. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addProperty()
		case 2:
			listProperties()
		case 3:
			addUnitToProperty()
		case 4:
			listUnitsForProperty()
		case 5:
			addLeaseToUnit()
		case 6:
			listLeasesForUnit()
		case 7:
			recordRentPayment()
		case 8:
			generateRentRoll()
		case 9:
			saveData()
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

// --- Property Management ---

func addProperty() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Add New Property ---")
	fmt.Print("Enter property name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter property address: ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)

	newProperty := Property{
		ID:      generateID(),
		Name:    name,
		Address: address,
		Units:   []Unit{},
	}

	properties = append(properties, newProperty)
	saveData()
	fmt.Println("✅ Property added successfully!\n")
}

func listProperties() {
	fmt.Println("\n--- List of Properties ---")

	if len(properties) == 0 {
		fmt.Println("No properties found.")
		return
	}

	for i, p := range properties {
		fmt.Printf("%d) %s - %s (ID: %s)\n", i+1, p.Name, p.Address, p.ID)
	}
	fmt.Println()
}

// --- Unit Management ---

func addUnitToProperty() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Add New Unit ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available. Please add a property first.\n")
		return
	}

	listProperties()

	var propChoice int
	fmt.Print("Select the property number to add a unit to: ")
	fmt.Scan(&propChoice)

	if propChoice < 1 || propChoice > len(properties) {
		fmt.Println("❌ Invalid property selection.\n")
		return
	}

	selectedProperty := &properties[propChoice-1]

	fmt.Print("Enter unit number (e.g., 2A): ")
	unitNumber, _ := reader.ReadString('\n')
	unitNumber = strings.TrimSpace(unitNumber)

	var bedrooms int
	fmt.Print("Enter number of bedrooms: ")
	fmt.Scan(&bedrooms)

	var bathrooms float64
	fmt.Print("Enter number of bathrooms: ")
	fmt.Scan(&bathrooms)

	var squareFeet int
	fmt.Print("Enter square footage: ")
	fmt.Scan(&squareFeet)

	newUnit := Unit{
		ID:         generateID(),
		UnitNumber: unitNumber,
		Bedrooms:   bedrooms,
		Bathrooms:  bathrooms,
		SquareFeet: squareFeet,
		Leases:     []Lease{},
	}

	selectedProperty.Units = append(selectedProperty.Units, newUnit)
	saveData()
	fmt.Println("✅ Unit added successfully!\n")
}

func listUnitsForProperty() {
	fmt.Println("\n--- List Units for a Property ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available. Please add a property first.\n")
		return
	}

	listProperties()

	var propChoice int
	fmt.Print("Select the property number to view units: ")
	fmt.Scan(&propChoice)

	if propChoice < 1 || propChoice > len(properties) {
		fmt.Println("❌ Invalid property selection.\n")
		return
	}

	selectedProperty := properties[propChoice-1]

	if len(selectedProperty.Units) == 0 {
		fmt.Println("No units found for this property.\n")
		return
	}

	fmt.Printf("\nUnits in %s (%s):\n", selectedProperty.Name, selectedProperty.Address)

	for i, unit := range selectedProperty.Units {
		fmt.Printf("%d) Unit: %s | Beds: %d | Baths: %.1f | Sqft: %d\n",
			i+1, unit.UnitNumber, unit.Bedrooms, unit.Bathrooms, unit.SquareFeet)
	}
	fmt.Println()
}

// --- Lease and Tenant Management ---

func addLeaseToUnit() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- Create Lease for a Unit ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available. Please add a property first.\n")
		return
	}

	listProperties()

	var propChoice int
	fmt.Print("Select the property number: ")
	fmt.Scan(&propChoice)

	if propChoice < 1 || propChoice > len(properties) {
		fmt.Println("❌ Invalid property selection.\n")
		return
	}

	selectedProperty := &properties[propChoice-1]

	if len(selectedProperty.Units) == 0 {
		fmt.Println("❌ No units available in this property. Please add a unit first.\n")
		return
	}

	fmt.Printf("\nUnits in %s (%s):\n", selectedProperty.Name, selectedProperty.Address)
	for i, unit := range selectedProperty.Units {
		fmt.Printf("%d) Unit: %s\n", i+1, unit.UnitNumber)
	}

	var unitChoice int
	fmt.Print("Select the unit number: ")
	fmt.Scan(&unitChoice)

	if unitChoice < 1 || unitChoice > len(selectedProperty.Units) {
		fmt.Println("❌ Invalid unit selection.\n")
		return
	}

	selectedUnit := &selectedProperty.Units[unitChoice-1]

	fmt.Println("\n--- Enter Tenant Information ---")
	fmt.Print("Tenant Full Name: ")
	tenantName, _ := reader.ReadString('\n')
	tenantName = strings.TrimSpace(tenantName)

	fmt.Print("Tenant Email: ")
	tenantEmail, _ := reader.ReadString('\n')
	tenantEmail = strings.TrimSpace(tenantEmail)

	fmt.Print("Tenant Phone: ")
	tenantPhone, _ := reader.ReadString('\n')
	tenantPhone = strings.TrimSpace(tenantPhone)

	tenant := Tenant{
		ID:       generateID(),
		FullName: tenantName,
		Email:    tenantEmail,
		Phone:    tenantPhone,
	}

	fmt.Println("\n--- Enter Lease Information ---")
	var startYear, startMonth, startDay int
	fmt.Print("Lease Start Date (YYYY MM DD): ")
	fmt.Scan(&startYear, &startMonth, &startDay)

	var endYear, endMonth, endDay int
	fmt.Print("Lease End Date (YYYY MM DD): ")
	fmt.Scan(&endYear, &endMonth, &endDay)

	var monthlyRent float64
	fmt.Print("Monthly Rent: ")
	fmt.Scan(&monthlyRent)

	var securityDeposit float64
	fmt.Print("Security Deposit: ")
	fmt.Scan(&securityDeposit)

	lease := Lease{
		ID:              generateID(),
		Tenant:          tenant,
		StartDate:       time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC),
		EndDate:         time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC),
		MonthlyRent:     monthlyRent,
		SecurityDeposit: securityDeposit,
		RentHistory:     []RentPayment{},
		Status:          "active",
	}

	selectedUnit.Leases = append(selectedUnit.Leases, lease)
	saveData()
	fmt.Println("✅ Lease created successfully!\n")
}

func listLeasesForUnit() {
	fmt.Println("\n--- List Leases for a Unit ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available.\n")
		return
	}

	listProperties()

	var propChoice int
	fmt.Print("Select the property number: ")
	fmt.Scan(&propChoice)

	if propChoice < 1 || propChoice > len(properties) {
		fmt.Println("❌ Invalid property selection.\n")
		return
	}

	selectedProperty := &properties[propChoice-1]

	if len(selectedProperty.Units) == 0 {
		fmt.Println("❌ No units available in this property.\n")
		return
	}

	fmt.Printf("\nUnits in %s (%s):\n", selectedProperty.Name, selectedProperty.Address)
	for i, unit := range selectedProperty.Units {
		fmt.Printf("%d) Unit: %s\n", i+1, unit.UnitNumber)
	}

	var unitChoice int
	fmt.Print("Select the unit number: ")
	fmt.Scan(&unitChoice)

	if unitChoice < 1 || unitChoice > len(selectedProperty.Units) {
		fmt.Println("❌ Invalid unit selection.\n")
		return
	}

	selectedUnit := selectedProperty.Units[unitChoice-1]

	if len(selectedUnit.Leases) == 0 {
		fmt.Println("❌ No leases found for this unit.\n")
		return
	}

	fmt.Printf("\nLeases for Unit %s:\n", selectedUnit.UnitNumber)

	for i, lease := range selectedUnit.Leases {
		fmt.Printf("%d) Tenant: %s | Start: %s | End: %s | Rent: $%.2f | Status: %s\n",
			i+1,
			lease.Tenant.FullName,
			lease.StartDate.Format("2006-01-02"),
			lease.EndDate.Format("2006-01-02"),
			lease.MonthlyRent,
			lease.Status,
		)
	}
	fmt.Println()
}

func recordRentPayment() {

	fmt.Println("\n--- Record Rent Payment ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available.\n")
		return
	}

	listProperties()

	var propChoice int
	fmt.Print("Select the property number: ")
	fmt.Scan(&propChoice)

	if propChoice < 1 || propChoice > len(properties) {
		fmt.Println("❌ Invalid property selection.\n")
		return
	}

	selectedProperty := &properties[propChoice-1]

	if len(selectedProperty.Units) == 0 {
		fmt.Println("❌ No units available.\n")
		return
	}

	fmt.Printf("\nUnits in %s (%s):\n", selectedProperty.Name, selectedProperty.Address)
	for i, unit := range selectedProperty.Units {
		fmt.Printf("%d) Unit: %s\n", i+1, unit.UnitNumber)
	}

	var unitChoice int
	fmt.Print("Select the unit number: ")
	fmt.Scan(&unitChoice)

	if unitChoice < 1 || unitChoice > len(selectedProperty.Units) {
		fmt.Println("❌ Invalid unit selection.\n")
		return
	}

	selectedUnit := &selectedProperty.Units[unitChoice-1]

	if len(selectedUnit.Leases) == 0 {
		fmt.Println("❌ No leases available for this unit.\n")
		return
	}

	fmt.Printf("\nLeases for Unit %s:\n", selectedUnit.UnitNumber)
	for i, lease := range selectedUnit.Leases {
		fmt.Printf("%d) Tenant: %s | Status: %s\n", i+1, lease.Tenant.FullName, lease.Status)
	}

	var leaseChoice int
	fmt.Print("Select the lease number: ")
	fmt.Scan(&leaseChoice)

	if leaseChoice < 1 || leaseChoice > len(selectedUnit.Leases) {
		fmt.Println("❌ Invalid lease selection.\n")
		return
	}

	selectedLease := &selectedUnit.Leases[leaseChoice-1]

	var year, month int
	fmt.Print("Enter payment year and month (YYYY MM): ")
	fmt.Scan(&year, &month)

	var amountPaid float64
	fmt.Print("Enter amount paid: ")
	fmt.Scan(&amountPaid)

	paymentDate := time.Now()

	rentPayment := RentPayment{
		Month:      time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC),
		AmountDue:  selectedLease.MonthlyRent,
		AmountPaid: amountPaid,
		PaidDate:   &paymentDate,
		Status:     calculatePaymentStatus(selectedLease.MonthlyRent, amountPaid),
		Notes:      "",
	}

	selectedLease.RentHistory = append(selectedLease.RentHistory, rentPayment)

	saveData()
	fmt.Println("✅ Rent payment recorded successfully!\n")
}

func generateRentRoll() {
	fmt.Println("\n--- Rent Roll Report ---")

	if len(properties) == 0 {
		fmt.Println("❌ No properties available.\n")
		return
	}

	var totalExpected float64
	var totalCollected float64
	var totalOutstanding float64

	for _, property := range properties {
		for _, unit := range property.Units {
			for _, lease := range unit.Leases {
				if lease.Status == "active" {
					fmt.Printf("\nProperty: %s | Unit: %s | Tenant: %s\n",
						property.Name, unit.UnitNumber, lease.Tenant.FullName)

					if len(lease.RentHistory) == 0 {
						fmt.Println("No payments recorded yet.")
						totalExpected += lease.MonthlyRent
						totalOutstanding += lease.MonthlyRent
						continue
					}

					for _, payment := range lease.RentHistory {
						statusIcon := getStatusIcon(payment.Status)
						fmt.Printf("  [%s] %s | Due: $%.2f | Paid: $%.2f\n",
							statusIcon,
							payment.Month.Format("Jan 2006"),
							payment.AmountDue,
							payment.AmountPaid)

						totalExpected += payment.AmountDue
						totalCollected += payment.AmountPaid

						if payment.AmountPaid < payment.AmountDue {
							totalOutstanding += payment.AmountDue - payment.AmountPaid
						}
					}
				}
			}
		}
	}

	fmt.Println("\n==== Rent Roll Summary ====")
	fmt.Printf("Total Rent Expected: $%.2f\n", totalExpected)
	fmt.Printf("Total Rent Collected: $%.2f\n", totalCollected)
	fmt.Printf("Total Outstanding: $%.2f\n", totalOutstanding)
	fmt.Println("===========================\n")
}

func getStatusIcon(status string) string {
	switch status {
	case "paid":
		return "✔️"
	case "partial":
		return "⚠️"
	case "unpaid":
		return "❌"
	default:
		return "❓"
	}
}

func calculatePaymentStatus(amountDue, amountPaid float64) string {
	if amountPaid >= amountDue {
		return "paid"
	} else if amountPaid > 0 {
		return "partial"
	}
	return "unpaid"
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()+rand.Int63n(1000))
}

func saveData() {
	data, err := json.MarshalIndent(properties, "", "  ")
	if err != nil {
		fmt.Println("❌ Failed to encode data:", err)
		return
	}

	err = ioutil.WriteFile(dataFile, data, 0644)
	if err != nil {
		fmt.Println("❌ Failed to write data:", err)
	} else {
		fmt.Println("💾 Data saved successfully.")
	}
}

func loadData() {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		fmt.Println("📂 No saved data found. Starting fresh.")
		return
	}

	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		fmt.Println("❌ Failed to read data:", err)
		return
	}

	err = json.Unmarshal(data, &properties)
	if err != nil {
		fmt.Println("❌ Failed to decode data:", err)
	} else {
		fmt.Println("✅ Loaded saved properties.")
	}
}
