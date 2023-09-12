package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shaPad(input string, blockSize int) string {
	// Append "1" to the input
	input += "1"

	// Calculate the number of "0"s needed to reach the desired length
	K := (blockSize - (len(input) + 8)) % blockSize

	// Append K "0"s to the input
	input += strings.Repeat("0", K)

	// Append the 64-bit block that represents the length of the original message
	length := uint64(len(input) - 1) // Length without the appended "1"
	lengthBinary := fmt.Sprintf("%064b", length)
	input += lengthBinary

	return input
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter binary input: ")
	originalMessage, _ := reader.ReadString('\n')
	originalMessage = strings.TrimSpace(originalMessage)

	blockSize := 512

	paddedMessage := shaPad(originalMessage, blockSize)

	fmt.Println("Original Input:", originalMessage)
	fmt.Println("Padded Output:", paddedMessage)
}
