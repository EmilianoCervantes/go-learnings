package main

import (
	"fmt"
	"sync"

	c "concurrency/channels"
)

func speak(message string, wg *sync.WaitGroup) {
	// defer is a statement that defers the execution of a function until the surrounding function returns
	// It is used to ensure that a clean-up is performed
	// In this context, it is used to ensure that the WaitGroup.Done() is called

	if wg != nil { // Avoiding panic: runtime error: invalid memory address or nil pointer dereference
		defer wg.Done()
	}

	fmt.Println(message)
}

// Main actually happens inside a goroutine
// Syntax: go <operation>
// More on https://go.dev/tour/concurrency/1
func BasicSyntax() {
	// speak("Hello") // Gets printed
	// If we leave it like this, we will notice "World" is never printed
	// Why?
	// As the operation is concurrent, after the first 'speak', the one without a 'go' finishes, main also finishes.
	// Basically, it is not part of the same execution
	// go speak("World") // WON'T get printed

	// We could do the NOT recommended
	// time.Sleep(time.Second * 1) // But this is not efficient and against what you are trying to achieve with concurrency

	// If you run the code above, only "Hello" gets printed
}

func MoreConcreteExample() {
	var waitGroup sync.WaitGroup // Accumulate goroutines

	fmt.Println("Hello") // Gets executed as per usual

	// Add a goroutine to the WaitGroup
	// Add(counter) - counter = number of goroutines that you will add
	// That way the program waits for all goroutines to finish
	waitGroup.Add(1)

	go speak("World", &waitGroup) // Now with the WaitGroup, main awaits for "World" to be printed

	waitGroup.Wait() // That way "main" waits for all goroutines to be done via wg.Done()
}

// It even works with anonymous functions
/**
 * You'll notice how prints can be ordered differently
 * Hello -> Baby -> Hasta la vista -> World
 * Hello -> Baby -> World -> Hasta la vista
 */
func main() {
	var waitGroup sync.WaitGroup // Accumulate goroutines

	speak("Hello", nil)

	waitGroup.Add(3)

	go speak("World", &waitGroup)

	go func() {
		defer waitGroup.Done()
		fmt.Println("Hasta la vista")
	}()

	go func(text string) {
		defer waitGroup.Done()
		fmt.Println(text)
	}("Baby")

	waitGroup.Wait()

	fmt.Println()

	c.Channels()

	c.MultipleChannels()
}
