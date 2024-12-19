package main

import "fmt"

const Skip = true

func conditionals() {
	if Skip {
		return
	}

	example1 := 1
	example2 := 2

	fmt.Println(example1 == example2)           // false
	fmt.Println(example1 != example2)           // true
	fmt.Println(example1 < example2)            // true
	fmt.Println(example1 <= example2)           // true
	fmt.Println(example1 > example2)            // false
	fmt.Println(example1 >= example2)           // false
	fmt.Println(1 == example1 && 1 == example2) // false
	fmt.Println(1 == example1 || 1 == example2) // true
}

func usingSwitch() {
	if Skip {
		return
	}

	remainder := 8961290 % 15

	switch remainder {
	case 0:
		fmt.Println("multiple")
	default:
		fmt.Println("The remainder is:", remainder)
	}
}

/**
 * Use cases for defer:
 * Closing a connection
 * Closing files
 */
func usingDefer() {
	if Skip {
		return
	}

	defer fmt.Println("Hola")
	fmt.Println("Mundo")
}

func continueAndBreak() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			// When something not uncommon and somewhat expected happens
			// And basically you want your code to continue
			continue
		}

		if i > 5 {
			fmt.Println("Limit reached at:", i)
			break
		}

		fmt.Println("Regular print:", i) // We will never print odd numbers
	}
}

func main() {
	conditionals()
	usingSwitch()
	usingDefer()
	continueAndBreak()
}
