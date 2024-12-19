package main

import (
	"fmt"
	"reflect"
)

func numbers() {
	/**
	 * We have int(8,16,32,64) and uint(8,16,32,64)
	 * float(32,64)
	 * complex(64,128) Real or imaginary float(32,64)
	 */

	fmt.Println("We have int:", reflect.TypeOf(8))

	var num2 uint64 = 12
	fmt.Println("uint:", reflect.TypeOf(num2))

	myFloat := 12.1
	fmt.Println("floats:", reflect.TypeOf(myFloat))

	var complex complex128 = 12.1
	fmt.Println("As well as complex:", reflect.TypeOf(complex))
}

func main() {
	numbers()
	fmt.Println("Strings:", reflect.TypeOf(""))    // string
	fmt.Printf("Strings: %T\n", "")                // string
	fmt.Println("Booleans:", reflect.TypeOf(true)) // bool

	fmt.Println()

	fmt.Println("Let's print every variable together")
	fmt.Printf("Here a %s, then a %d and also a %f\n\n", "string example", 1, 1.1)
	fmt.Println("If you try to add to the print the wrong type, it will throw an error like '%!d' or '%!f' with '!' saying: hey, this is NOT the type.")

	fmt.Println()

	fmt.Println("But as always, we can be lazy and just do '%v' and not care about the type")
	fmt.Printf("%v, %v, %v\n", "string example", 1, 1.1)
}
