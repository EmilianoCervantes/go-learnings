package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const Skip = true

func readAndTrimLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("There was an error with the input message:", err)
		os.Exit(1)
	}

	if len(line) < 2 {
		log.Fatal("The input cannot be empty:")
		os.Exit(1)
	}

	return strings.TrimSpace(line)
}

func myFirstLoop(message string, repetitions int) {
	if Skip {
		return
	}

	for i := 1; i <= repetitions; i++ {
		res := fmt.Sprintf("%s for the %dth time", message, i)
		fmt.Println(res)
	}
}

func forWhile(message string, repetitions int) {
	if Skip {
		return
	}

	count := 0

	for count < repetitions {
		count++
		res := fmt.Sprintf("%s for the %dth time", message, count)
		fmt.Println(res)
	}
}

func forForever() {
	if Skip {
		return
	}

	count := 0

	for {
		fmt.Println("Counter: ", count)
		count++
	}
}

func main() {
	message := readAndTrimLine("Add a message to repeat: ")
	repetitions := readAndTrimLine("How many times do you want to repeat it? ")

	repetitionsAsInt, err := strconv.Atoi(repetitions)

	if err != nil {
		log.Fatal("This is not a number: ", err)
		os.Exit(1)
	}

	myFirstLoop(message, repetitionsAsInt)
	forWhile(message, repetitionsAsInt)
	forForever()
}
