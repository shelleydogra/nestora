package main

import (
    "time"
)

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
    Month       time.Time
    AmountDue   float64
    AmountPaid  float64
    PaidDate    *time.Time
    Status      string // paid, unpaid, partial
    Notes       string
}

func main() {
    // We will build CLI commands here next (add property, add unit, add tenant, add lease, etc.)
}
