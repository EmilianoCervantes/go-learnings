package main

import "fmt"

func InterfacesSlice() {
	// While Go is strict with its types, there a trick to do a slice of multiple kinds of data
	mySlice := []interface{}{"Hello", 12345, 104.5}
	mySlice = append(mySlice, "World")

	fmt.Println(mySlice)

	mySlice[0] = 3.1415
	mySlice[2] = "interfaces"
	mySlice[3] = []rune{1}

	fmt.Println(mySlice)
}
