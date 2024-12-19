package main

import "fmt"

func addVsSubtract(_ string, num1, num2 int) (res1, res2 int) {
	sum := num1 + num2
	subs := num1 - num2

	return sum, subs
}

func main() {
	// assignment mismatch: 1 variable but addVsSubtract returns 2 values compiler(WrongAssignCount)
	// sum := addVsSubtract("Example", 10, 20)

	// But I don't want to use both
	sum, _ := addVsSubtract("Example", 10, 20)
	_, subs := addVsSubtract("Example", 10, 20)

	fmt.Println("sum:", sum)
	fmt.Println("subs:", subs)
}
