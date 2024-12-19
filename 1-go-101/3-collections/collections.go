package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const Skip = true

/**
* Arrays in Go
* - Arrays are fixed-size sequences of elements of the same type.
* - The size of an array is part of its type.
* - Arrays are value types.
* - Arrays are declared using the following syntax:
*   var <arrayName> [<size>]<type>
* - Arrays can be initialized using the following syntax:
*   <arrayName> := [<size>]<type>{<comma-separated values>}
 */
func arraysInGo() {
	if Skip {
		return
	}

	var array [4]int
	fmt.Println(("Array initialized with default values:"), array)

	array[0] = 1
	array[1] = 2
	fmt.Println("My first array:", array, "with length:", len(array), "and capacity:", cap(array))

	fmt.Println("First element of the slice:", array[0])
	fmt.Println("Last element of the array:", array[len(array)-1])
	// [startIndex:endIndex]
	// [inclusive:exclusive]
	fmt.Println("Array from index 2 to 3 [2:3]:", array[2:3])
	fmt.Println("Array from index 2 to 3 [2:4]:", array[2:4])
	// [startIndex:]
	// [inclusive:end]
	fmt.Println("Array from index 3 to the end [3:]:", array[3:])
	fmt.Println("Array from index 4 to the end [4:]:", array[4:])
	// [:endIndex]
	// [beginning:exclusive]
	fmt.Println("Array from the beginning to index 2 [:3]:", array[:3])
	fmt.Println()
}

/**
 * Slices in Go
 * - Slices are dynamic arrays.
 * - Slices are references to arrays.
 * - Slices are declared using the following syntax:
 *   var <sliceName>[]<type>
 * - Slices can be initialized using the following syntax:
 *   <sliceName> := []<type>{<comma-separated values>}
 * Difference between arrays and slices:
 * - Arrays are fixed-size sequences of elements of the same type.
 * - Slices are dynamic arrays.
 */
func slicesInGo() {
	if Skip {
		return
	}

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("Slice initialized:", slice, "with length:", len(slice), "and capacity:", cap(slice))
	fmt.Println()

	/** Methods for slices
	 * - append(slice, element) -> appends a new element to the slice.
	 * - copy(destination, source) -> copies elements from the source slice to the destination slice.
	 * - len(slice) -> returns the length of the slice.
	 * - cap(slice) -> returns the capacity of the slice.
	 */

	slice = append(slice, 7)
	fmt.Println("Slice after appending a new element:", slice, "with length:", len(slice), "and capacity:", cap(slice))

	// Append new list
	newSlice := []int{8, 9, 10}
	// ... -> unpacks the slice
	// Similar to JavaScript's spread operator
	// JavaScript: [...slice, ...newSlice] and in Go: append(slice, newSlice...)
	slice = append(slice, newSlice...)
}

func printNewLineIfValueIsTheLastElement(slice []string, value string) {
	if value == slice[len(slice)-1] {
		fmt.Println()
	}
}

/**
 * Iterating over arrays and slices
 * - Arrays and slices can be iterated using the for loop.
 * - The range keyword is used to iterate over arrays and slices.
 * - The range keyword returns the index and the value of the element at that index.
 */
func iteratingSlices() {
	if Skip {
		return
	}

	slice := []string{"hello", "world", "from", "Go", "!"}

	for index, value := range slice {
		fmt.Println("Index:", index, "Value:", value)
		printNewLineIfValueIsTheLastElement(slice, value)
	}

	// If you don't need the index, you can use the blank identifier _
	fmt.Println("Getting only the values:")
	for _, value := range slice {
		fmt.Println("Value:", value)
		printNewLineIfValueIsTheLastElement(slice, value)
	}

	// And with the value, you can ask only for the index
	fmt.Println("Getting only the indexes:")

	for index := range slice {
		fmt.Println("Index:", index)
		if index == len(slice)-1 {
			fmt.Println()
		}
	}
}

func readAndTrimLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("There was an error with the input message:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(line)
}

func isPalindrome() {
	if Skip {
		return
	}

	message := readAndTrimLine("Give me a potential palindrome: ")
	option := readAndTrimLine("Do you want to check for case sensitivity? (Y/n): ")

	if option != "Y" && option != "y" && option != "" && option != "n" && option != "N" {
		log.Fatal("Invalid option. Please provide a valid option.")
		os.Exit(1)
	}

	// Convert the string to a slice of letters
	// A rune is a type alias for int32 and is used to represent a Unicode code point.
	// A code point is a unique integer value that represents a single unique character.
	if option == "N" || option == "n" {
		message = strings.ToLower(message)
	}

	letters := []rune(message)
	// Iterate over the letters in the slice
	for i := 0; i < len(letters)/2; i++ {
		// Compare the letters at the beginning and the end of the slice
		if letters[i] != letters[len(letters)-1-i] {
			// If the letters are different, the text is not a palindrome
			fmt.Printf("Is the text '%s' a palindrome? %t\n", message, false)
		}
	}

	fmt.Printf("Is the text '%s' a palindrome? %t\n", message, true)
}

/**
 * Maps in Go
 * - Maps are key-value pairs.
 * - Maps are unordered collections.
 * - Maps are declared using the following syntax:
 *   var <mapName> map[<keyType>]<valueType>
 * - Maps can be initialized using the following syntax:
 *   <mapName> := map[<keyType>]<valueType>{<key>: <value>, <key>: <value>}
 * We can also use the make function to create a map:
 *   <mapName> := make(map[<keyType>]<valueType>)
 * Maps are like dictionaries in Python.
 * Maps are like hash tables in other programming languages.
 */
func mapsInGo() {
	if Skip {
		return
	}

	fmt.Println("Using map[<keyType>]<valueType>{<key>: <value>, <key>: <value>}")
	myMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("My first map:", myMap, "with length:", len(myMap))

	// Adding a new key-value pair
	myMap["four"] = 4
	fmt.Println("My map after adding a new key-value pair:", myMap, "with length:", len(myMap))

	// Deleting a key-value pair
	delete(myMap, "two")
	fmt.Println("My map after deleting a key-value pair:", myMap)
	fmt.Println()

	fmt.Println("Using make(map[<keyType>]<valueType>)")
	agesMap := make(map[string]int)
	fmt.Println("Another map with make:", agesMap, "with length:", len(agesMap))

	agesMap["Alice"] = 25
	agesMap["Bob"] = 30

	fmt.Println("My map with make after adding two ages:", agesMap, "with length:", len(agesMap))

	// Iterating over maps
	for key, value := range agesMap {
		fmt.Println("Person:", key, "- Age:", value)
	}

	// Remove a key-value pair
	delete(agesMap, "Alice")

	fmt.Println("My map with make after deleting a key-value pair:", agesMap, "with length:", len(agesMap))

	// Adding a new key-value pair
	agesMap["Charlie"] = 35
	agesMap["David"] = 40

	// Check if a key exists in the map
	value, ok := agesMap["Charlie"]
	fmt.Println("Is Charlie in the map?", ok, "with value:", value)

	// This is how you would actually use this in your code
	if _, ok := agesMap["Charlie"]; ok {
		fmt.Println("Charlie is in the map - age:", agesMap["Charlie"])
	}
}

/**
 * Structs in Go
 * - Structs are user-defined types.
 * - Structs are collections of fields.
 * - Structs are declared using the following syntax:
 *   type <structName> struct {
 *     <fieldName> <fieldType>
 *     <fieldName> <fieldType>
 *     ...
 *   }
 * - Structs can be initialized using the following syntax:
 *   <structName>{<fieldValue>, <fieldValue>, ...}
 * Structs are like classes in other programming languages.
 */

type car struct {
	make  string
	model string
	year  int
}

func structs() {
	if Skip {
		return
	}

	// Fields provide the values in the order they are declared
	myCar := car{"Mazda", "2", 2021}
	fmt.Println("My car:", myCar)
	fmt.Println("My car's model:", myCar.model)
	fmt.Println()

	// We don't need to provide all the fields
	// We also don't need to provide the fields in order
	anotherCar := car{year: 2020, make: "Audi"}
	fmt.Println("Another car:", anotherCar)
	fmt.Println()

	// Empty struct
	var emptyCar car
	fmt.Println("Empty car:", emptyCar)
}

func main() {
	arraysInGo()
	slicesInGo()
	iteratingSlices()
	isPalindrome()
	mapsInGo()
	structs()
}
