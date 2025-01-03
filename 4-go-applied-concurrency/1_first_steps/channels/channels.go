package channelsExample

import "fmt"

const skip = true

// chan is the reserved keyword to create channels
// A good practice is to indicate id a channel will be used only for adding data or extracting it
// Also, it helps performance
// So, from:
// func shout(message string, channel chan string) {
// Instead lets do
func shout(message string, channel chan<- string) {
	// "<-" means adding something to the channel
	channel <- message
}

/**
 * Channels in Go
 * They are the way to order goroutines
 * Channels are the pipes that connect concurrent goroutines
 * You can send values into channels from one goroutine and receive those values into another goroutine
 * Channels are typed
 */
func Channels() {
	if skip {
		return
	}

	// Syntax: make(chan <type>, <optional: number of simultaneous data the channel will manage>)
	// If optional is removed, then the number of data that can be manage will be dynamic
	// While it is optional, it is recommended for performance
	// chan is the reserved keyword to create channels
	// chan <type>
	channel := make(chan string, 1)

	fmt.Println("Hello")

	go shout("Bye", channel)

	// If channel <- means adding, with "<-" to the right
	// "<-" to the left means extracting data from the channel
	fmt.Println(<-channel)
}
