package main

import (
	"fmt"
	"math"

	// Importing from another package
	// aliasName "packagePath"
	vehicles "packageExample/packageExample"
	"reflect"
)

const Skip = true

func variables101() {
	if Skip {
		return
	}

	age := 15

	fmt.Println(age)
}

func math101() {
	if Skip {
		return
	}

	pi := 3.1416

	fmt.Println(reflect.TypeOf(pi))

	radius := 14
	var area float64

	fmt.Println("Area initialized:", area) // Area initialized: 0

	area = pi * math.Pow(float64(radius), 2)

	fmt.Println("Area result:", area)                 // 615.7536
	fmt.Printf("Area 1 decimal place: %.1f\n", area)  // 615.8
	fmt.Printf("Area 0 decimal places: %.0f\n", area) // 616
}

func math102() {
	if Skip {
		return
	}

	myNumber := 100
	fmt.Println("We have any number:", myNumber)

	myNumber++
	fmt.Println("We can do myNumber++", myNumber)
	// fmt.Println("BUT not ++myNumber", ++myNumber) // expected operand, found '++'
	fmt.Println("BUT not ++myNumber: expected operand, found '++'") // expected operand, found '++'
}

// Instead of "num1 int, num2 int"
func division(num1, num2 int) {
	if Skip {
		return
	}

	result := num1 / num2

	fmt.Println("Divide", num1, "/", num2)
	fmt.Println("Division result:", result)
	fmt.Println()
	fmt.Println("Why 33? Where are the decimals?")
	fmt.Println()

	message := fmt.Sprintf("Divide float64(%d) / float64(%d)", num1, num2)
	fmt.Println(message)
	result2 := float64(num1) / float64(num2)
	fmt.Println("Division result #2:", result2)
	fmt.Println()

	fmt.Println("Remanent", num1, "%", num2)
	result = num1 % num2
	fmt.Println("Remanent result:", result)
}

func publicPrivateExample() {
	if Skip {
		return
	}

	vehicles.SeePotentialVehicles()

	var myBoat vehicles.Boat
	fmt.Println("My boat:", myBoat)
}

func main() {
	variables101()
	math101()
	math102()
	division(100, 3)
	publicPrivateExample()
}
