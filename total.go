package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var userInput int
	fmt.Print("Enter the number of zeros you want: ")
	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	if userInput < 0 {
		fmt.Println("Number of zeros cannot be negative.")
		return
	}

	zeros := strings.Repeat("0", userInput)

	// Call the binary_hasher program with zeros as a command-line argument
	hashBinaryInput(zeros)
}

func hashBinaryInput(binaryInput string) {
	// Create a new process to run the binary_hasher program
	cmd := exec.Command("./test", binaryInput)

	// Set the output to the current terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the process
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
