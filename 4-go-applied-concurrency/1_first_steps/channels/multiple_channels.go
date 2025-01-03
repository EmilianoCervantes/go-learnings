package channelsExample

import (
	"fmt"
)

func channelStatus(channel chan string) {
	fmt.Println("Channel Status -> Channel length/# of messages:", len(channel), "- Channel, capacity:", cap(channel))
}

func emailMessages(email string, channel chan<- string) {
	channel <- email
}

func MultipleChannels() {
	if !skip {
		return
	}

	channel := make(chan string, 3)
	channel <- "Message 1"
	channel <- "Message 2"
	// If we were to send a message OVER the *capacity*,
	// we would get -> fatal error: all goroutines are asleep - deadlock!
	// channel <- "Message #4"

	// len(channel) - will tell us how many goroutines/data are inside a channel
	// cap(channel) - will return the MAX amount of goroutines the channel can store
	channelStatus(channel) // 2 - 3

	// RANGE & CLOSE

	fmt.Println("----- CHANNEL \"CLOSE\" -----")

	/**
	 * close(channel) - will close the channel
	 * This means the channel WON'T receive any new messages
	 * even when the channel might have capacity.
	 */
	close(channel)
	// channel <- "Message 3" // panic: send on closed channel

	/**
	 * for range channel
	 * will enable us to iterate through a channel.
	 *
	 * In this case the sintaxis is a little different from other "for index, value range string/array/slice {}"
	 * The range in channels "permits only one iteration variable".
	 * for message := range channel
	 */
	fmt.Println("----- CHANNEL \"RANGE\" -----")
	for message := range channel {
		fmt.Println(message)
	}

	channelStatus(channel) // 0 - 3

	fmt.Println()

	/**
	 * SELECT
	 */
	fmt.Println("----- CHANNEL \"SELECT\" -----")

	emailChannel1 := make(chan string)
	emailChannel2 := make(chan string)

	// Let's run some goroutines
	go emailMessages("Email 1", emailChannel1)
	go emailMessages("Email 2", emailChannel2)

	// At this point we don't know which goroutine will finish first

	// We will use select for this

	// In this example we are handling 2 channels
	numOfChannels := 2
	for i := 0; i < numOfChannels; i++ {
		select {
		case message1 := <-emailChannel1:
			fmt.Println("Email received from email 1:", message1)
		case message2 := <-emailChannel2:
			fmt.Println("Email received from email 2:", message2)
		}
	}
}
