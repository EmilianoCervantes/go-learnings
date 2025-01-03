package main

import "fmt"

type computer struct {
	brand string
	disk  int
	ram   int
}

/**
 * Pointer receiver
 * func (name pointer_receiver) methodName()
 * - Methods with a "pointer receiver" can modify the value to which the receiver points.
 * But only inside the method.
 * More in https://go.dev/tour/methods/4
 */
func (aPC computer) ping() {
	fmt.Println("Ping")
	fmt.Println("...pong", aPC.brand)
}

func (aPC computer) duplicateRAM() {
	fmt.Println("Before RAM:", aPC.ram)
	aPC.ram = aPC.ram * 2
	fmt.Println("After RAM:", aPC.ram)
}

/**
 * Value receiver
 * func (name *value_receiver) methodName()
 * Methods with a "*value_receiver" update the value of the receiver outside the method.
 */
func (aPC *computer) changeDisk() {
	aPC.disk = aPC.disk * 3
	fmt.Println("Inside changeDisk:", aPC.disk)
}

/**
 * String()
 * Is a method that returns a string representation of the object.
 * Just by implementing the String() method, when the object is printed, it will posses this format automatically.
 */
func (aPC computer) String() string {
	return fmt.Sprintf("My computer %d GB of RAM, %d GB in disk and is a %s\n", aPC.ram, aPC.disk, aPC.brand)
}

/**
 * Pointers in Go
 * - A pointer is a variable that stores the memory address of another variable.
 * - The type *T is a pointer to a T value.
 * - The zero value of a pointer is nil.
 * - The & operator generates a pointer to its operand.
 * - What & is doing behind the scenes is returning the memory address of the variable.
 * - The * operator denotes the pointer's underlying value.
 * - What * is doing behind the scenes is returning the value stored at the memory address = dereferencing.
 * - Again Dereferencing a pointer = accessing the value stored at the memory address pointed to by that pointer
 */
func main() {
	num1 := 50
	num2 := &num1
	num3 := *num2
	text := "Some text"
	boolTest := true

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	fmt.Println("*num2:", *num2)
	fmt.Println("&num1:", &num1)

	fmt.Println()
	fmt.Println("So I can change the value of the first variable through the second one:")
	fmt.Println("num1:", num1)
	fmt.Println("num2 := &num1:", *num2)
	fmt.Println("num3 := *num2", num3)
	fmt.Println("*num2 = 100")
	*num2 = 100
	fmt.Println("*num2:", *num2)
	fmt.Println("num1:", num1)
	fmt.Println("num3:", num3)

	fmt.Println()
	fmt.Println("We can get the memory address of variables")
	fmt.Println("text:", text)
	fmt.Println("&text:", &text)
	fmt.Println("&boolTest:", &boolTest)

	fmt.Println()
	fmt.Println("")
	myPC := computer{brand: "Lenovo", disk: 200, ram: 16}
	fmt.Println(myPC)
	myPC.ping()
	myPC.duplicateRAM()
	fmt.Println("Outside:", myPC.ram)
	myPC.changeDisk()
	fmt.Println("Outside changeDisk:", myPC.disk)
	fmt.Println("Finally, my PC:", myPC)

	fmt.Println()
}
