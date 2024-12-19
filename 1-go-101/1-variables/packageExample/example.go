package aPackageExample

import "fmt"

// Car is private as it starts with a lowercase letter
// type car struct {
// 	make  string // Brand in cars
// 	model string
// 	year  int
// }

// Boat is public as it starts with an uppercase letter
type Boat struct {
	Make  string
	Model string
	// year  int // "year" is private
	Year int // "Year" is public
}

// SeePotentialVehicles is public as it starts with an uppercase letter
func SeePotentialVehicles() {
	fmt.Println("See potential vehicles you can create:")
	fmt.Println("1. Boat")
}
